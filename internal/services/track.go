package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type TrackService struct {
	trackRepository *database.TrackRepository
}

func (t *TrackService) Create(track *models.Track) *gorm.DB {
	return t.trackRepository.Create(track)
}
