package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type TrackService struct {
	TrackRepository *database.TrackRepository
}

func (t *TrackService) Create(track *models.Track) *gorm.DB {
	return t.TrackRepository.Create(track)
}

func (t *TrackService) GetById(track *models.Track, id string) *gorm.DB {
	return t.TrackRepository.GetById(track, id)
}
