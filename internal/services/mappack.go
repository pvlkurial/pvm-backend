package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type MappackService struct {
	MappackRepository *database.MappackRepository
}

func (t *MappackService) Create(mappack *models.Mappack) *gorm.DB {
	return t.MappackRepository.Create(mappack)
}
