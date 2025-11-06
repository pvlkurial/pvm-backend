package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"
)

type PlayerService interface {
	Create(player *models.Player) error
	GetAll(players *[]models.Player) ([]models.Player, error)
	GetById(id string) (models.Player, error)
	Update(player *models.Player) error
}

type playerService struct {
	playerRepository database.PlayerRepository
}

func NewPlayerService(repo database.PlayerRepository) PlayerService {
	return &playerService{playerRepository: repo}
}

func (t *playerService) Create(player *models.Player) error {
	return t.playerRepository.Create(player)
}
func (t *playerService) GetAll(players *[]models.Player) ([]models.Player, error) {
	return t.playerRepository.GetAll()
}

func (t *playerService) GetById(id string) (models.Player, error) {
	return t.playerRepository.GetById(id)
}

func (t *playerService) Update(player *models.Player) error {
	return t.playerRepository.Update(player)
}
