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

func (t *PlayerRepository) GetAll(players *[]models.Player) *gorm.DB {
	return t.DB.Find(players)
}

func (t *PlayerRepository) GetById(player *models.Player, id string) *gorm.DB {
	return t.DB.Where("ID = ?", id).First(player)
}

func (t *PlayerRepository) Update(player *models.Player) *gorm.DB {
	return t.DB.Save(player)
}
