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

func (t *PlayerService) GetById(player *models.Player, id string) *gorm.DB {
	return t.PlayerRepository.GetById(player, id)
}

func (t *PlayerService) Update(player *models.Player) *gorm.DB {
	return t.PlayerRepository.Update(player)
}
