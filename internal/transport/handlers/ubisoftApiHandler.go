package handlers

import (
	"bytes"
	"encoding/json"
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/utils"
	"fmt"
	"net/http"
	"os"
)

type TokenManager struct {
	AccessToken      string
	RefreshToken     string
	AccessExpiresIn  int
	RefreshExpiresIn int
}

func fetchNadeoToken(audience string) (string, string, error) {

	body := map[string]string{
		"audience": audience,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	req, err := http.NewRequest("POST", utils.NadeoTokenURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	req.SetBasicAuth(os.Getenv("NADEO_API_USERNAME"), os.Getenv("NADEO_API_PASSWORD"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", os.Getenv("USER_AGENT"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch track: status %d\n", res.StatusCode)
		return "", "", err
	}
	var tokenResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
	json.NewDecoder(res.Body).Decode(&tokenResponse)
	return tokenResponse.AccessToken, tokenResponse.RefreshToken, nil
}

func (t *TokenManager) GetToken() string {
	if t.AccessToken == "" {
		fmt.Println("FETCHING NEW TOKEN")
		a, r, err := fetchNadeoToken("NadeoServices")
		t.AccessToken = a
		t.RefreshToken = r
		if err != nil {
			fmt.Print(err)
			return ""
		}
	}
	return t.AccessToken
}

func (t *TokenManager) FetchTrackInfo(trackid string) *models.Track {
	token := t.GetToken()
	req, err := http.NewRequest("GET", "https://prod.trackmania.core.nadeo.online/maps/"+trackid, nil)
	fmt.Println("DEBUG FOR URL: " + "https://prod.trackmania.core.nadeo.online/maps/" + trackid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Authorization", "nadeo_v1 t="+token)
	//req.Header.Add("User-Agent", os.Getenv("USER_AGENT"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch track: status %d\n", resp.StatusCode)
		return nil
	}

	track := &models.Track{}
	if err := json.NewDecoder(resp.Body).Decode(track); err != nil {
		fmt.Println(err)
		return nil
	}

	track.ID = track.MapID
	return track
}
