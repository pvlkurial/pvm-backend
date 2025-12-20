package clients

import (
	"example/pvm-backend/internal/models"
	"net/http"
	"time"
)

type TrackmaniaAPIClient struct {
	client       *http.Client
	accessToken  string
	refreshToken string
	expiresIn    string
}

func NewTrackmaniaAPIClient() *TrackmaniaAPIClient {
	return &TrackmaniaAPIClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (t *TrackmaniaAPIClient) FetchPlayerNames(playerIds []string) ([]models.Player, error) {
	return nil, nil
}
