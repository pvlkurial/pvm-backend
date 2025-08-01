package database

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type TrackRepository struct {
	DB *gorm.DB
}

func (t *TrackRepository) Create(track *models.Track) *gorm.DB {
	return t.DB.Create(&track)
}

func (t *TrackRepository) GetById(track *models.Track, id string) *gorm.DB {
	return t.DB.First(track).Where("ID = ?", id)
}
