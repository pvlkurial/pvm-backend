package repositories

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type PlayerRepository interface {
	Create(player *models.Player) error
	GetAll() ([]models.Player, error)
	GetById(id string) (models.Player, error)
	Update(player *models.Player) error
}

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) PlayerRepository {
	return &playerRepository{db: db}
}

func (t *playerRepository) Create(player *models.Player) error {
	err := t.db.Create(&player).Error
	return err
}

func (t *playerRepository) GetAll() ([]models.Player, error) {
	players := []models.Player{}
	err := t.db.Find(&players).Error
	return players, err
}

func (t *playerRepository) GetById(id string) (models.Player, error) {
	player := models.Player{}
	err := t.db.Where("ID = ?", id).First(&player).Error
	return player, err
}

func (t *playerRepository) Update(player *models.Player) error {
	err := t.db.Save(player).Error
	return err
}
