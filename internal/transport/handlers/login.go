package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/transport/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

const (
	TrackmaniaAuthURL  = "https://api.trackmania.com/oauth/authorize"
	TrackmaniaTokenURL = "https://api.trackmania.com/api/access_token"
	TrackmaniaUserURL  = "https://api.trackmania.com/api/user"
)

var stateStore = make(map[string]time.Time)

func cleanupExpiredStates() {
	for state, timestamp := range stateStore {
		if time.Since(timestamp) > 10*time.Minute {
			delete(stateStore, state)
		}
	}
}

func generateStateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func HandleTrackmaniaLogin(oauthConfig OAuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		state, err := generateStateToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate state token"})
			return
		}

		stateStore[state] = time.Now()

		params := url.Values{}
		params.Add("client_id", oauthConfig.ClientID)
		params.Add("redirect_uri", oauthConfig.RedirectURI)
		params.Add("response_type", "code")
		params.Add("state", state)

		authURL := fmt.Sprintf("%s?%s", TrackmaniaAuthURL, params.Encode())

		c.Redirect(http.StatusTemporaryRedirect, authURL)
	}
}

type TrackmaniaTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type TrackmaniaUserInfo struct {
	AccountID   string `json:"account_id"`
	DisplayName string `json:"display_name"`
	Username    string `json:"username,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
}

func HandleTrackmaniaCallback(db *gorm.DB, oauthConfig OAuthConfig, jwtConfig middleware.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		state := c.Query("state")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not provided"})
			return
		}

		if state == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "State parameter missing"})
			return
		}

		timestamp, exists := stateStore[state]
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state token"})
			return
		}

		if time.Since(timestamp) > 10*time.Minute {
			delete(stateStore, state)
			c.JSON(http.StatusBadRequest, gin.H{"error": "State token expired"})
			return
		}
		delete(stateStore, state)

		tokenResp, err := exchangeCodeForToken(oauthConfig, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token: " + err.Error()})
			return
		}

		userInfo, err := fetchTrackmaniaUserInfo(tokenResp.AccessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info: " + err.Error()})
			return
		}

		var user models.User
		result := db.Where("trackmania_account_id = ?", userInfo.AccountID).First(&user)

		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				now := time.Now()
				user = models.User{
					TrackmaniaAccountID:   userInfo.AccountID,
					TrackmaniaDisplayName: userInfo.DisplayName,
					TrackmaniaUsername:    userInfo.Username,
					AvatarURL:             userInfo.AvatarURL,
					AccessToken:           tokenResp.AccessToken,
					RefreshToken:          tokenResp.RefreshToken,
					TokenExpiry:           time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second),
					Role:                  models.RoleUser,
					LastLoginAt:           &now,
				}

				if err := db.Create(&user).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
					return
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
				return
			}
		} else {
			now := time.Now()
			user.TrackmaniaDisplayName = userInfo.DisplayName
			user.AccessToken = tokenResp.AccessToken
			user.RefreshToken = tokenResp.RefreshToken
			user.TokenExpiry = now.Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
			user.LastLoginAt = &now
			user.AvatarURL = userInfo.AvatarURL
			user.TrackmaniaUsername = userInfo.Username

			if err := db.Save(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
				return
			}
		}
		token, err := middleware.GenerateToken(jwtConfig, user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		frontendURL := os.Getenv("FRONTEND_URL")
		if frontendURL == "" {
			frontendURL = "http://localhost:3000"
		}

		redirectURL := fmt.Sprintf("%s/auth/callback?token=%s", frontendURL, token)
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}

func exchangeCodeForToken(config OAuthConfig, code string) (*TrackmaniaTokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", config.ClientID)
	data.Set("client_secret", config.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", config.RedirectURI)

	req, err := http.NewRequest("POST", TrackmaniaTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("token exchange failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp TrackmaniaTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func fetchTrackmaniaUserInfo(accessToken string) (*TrackmaniaUserInfo, error) {
	req, err := http.NewRequest("GET", TrackmaniaUserURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch user info with status %d: %s", resp.StatusCode, string(body))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println("Trackmania User API Response:", string(body))

	var userInfo TrackmaniaUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func HandleGetProfile(c *gin.Context) {
	user, err := middleware.GetUserFromGinContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":                      user.ID,
			"trackmania_account_id":   user.TrackmaniaAccountID,
			"trackmania_display_name": user.TrackmaniaDisplayName,
			"role":                    user.Role,
		},
	})
}
