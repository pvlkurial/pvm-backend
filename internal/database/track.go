package database

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type TrackRepository struct {
	db *gorm.DB
}

func (t *TrackRepository) Create(track *models.Track) *gorm.DB {
	return t.db.Create(&track)
}
