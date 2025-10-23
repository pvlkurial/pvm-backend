package middleware

import (
	"errors"
	"example/pvm-backend/internal/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
	ErrNoToken      = errors.New("no token provided")
	ErrUnauthorized = errors.New("unauthorized")
)

type contextKey string

const UserContextKey contextKey = "user"

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey     string
	TokenDuration time.Duration
}

func ValidateToken(config JWTConfig, tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
func GenerateToken(config JWTConfig, userID uint) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.TokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey))
}

func RequireAuth(db *gorm.DB, config JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractTokenFromGinContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
			c.Abort()
			return
		}
		claims, err := ValidateToken(config, tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid or expired token"})
			c.Abort()
			return
		}

		var user models.User
		if err := db.First(&user, claims.UserID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: user not found"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		c.Set("user", &user)
		c.Next()
	}
}
func extractTokenFromGinContext(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", ErrNoToken
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", ErrInvalidToken
	}

	return parts[1], nil
}

func GetUserFromGinContext(c *gin.Context) (*models.User, error) {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil, ErrUnauthorized
	}

	user, ok := userInterface.(*models.User)
	if !ok || user == nil {
		return nil, ErrUnauthorized
	}

	return user, nil
}

func RequireRole(role models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetUserFromGinContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if !user.HasRole(role) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func RequireAdmin() gin.HandlerFunc {
	return RequireRole(models.RoleAdmin)
}

func RequireContentCreator() gin.HandlerFunc {
	return RequireRole(models.RoleContentCreator)
}
