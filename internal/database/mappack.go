package database

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type MappackRepository struct {
	DB *gorm.DB
}

func (t *MappackRepository) Create(mappack *models.Mappack) *gorm.DB {
	return t.DB.Create(&mappack)
}
