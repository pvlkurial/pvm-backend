package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type PlayerService struct {
	PlayerRepository *database.PlayerRepository
}

func (t *PlayerService) Create(player *models.Player) *gorm.DB {
	return t.PlayerRepository.Create(player)
}
func (t *PlayerService) GetAll(players *[]models.Player) *gorm.DB {
	return t.PlayerRepository.GetAll(players)
}
