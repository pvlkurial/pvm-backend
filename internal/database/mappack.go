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
func (t *MappackRepository) GetById(mappack *models.Mappack, id string) *gorm.DB {
	return t.DB.First(mappack).Where("ID = ?", id)
}

func (t *MappackRepository) GetAll(mappacks *[]models.Mappack) *gorm.DB {
	return t.DB.Find(mappacks)
}
