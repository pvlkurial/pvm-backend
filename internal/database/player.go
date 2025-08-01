package database

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type PlayerRepository struct {
	DB *gorm.DB
}

func (t *PlayerRepository) Create(player *models.Player) *gorm.DB {
	return t.DB.Create(&player)
}
