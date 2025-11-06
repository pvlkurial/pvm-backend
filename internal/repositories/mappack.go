package repositories

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type MappackRepository interface {
	Create(mappack *models.Mappack) error
	GetById(id string) (models.Mappack, error)
	GetAll() ([]models.Mappack, error)
	CreateMappackTimeGoal(timegoal *models.TimeGoal) error
	GetAllMappackTimeGoals(mappackId string) ([]models.TimeGoal, error)
	RemoveTimeGoalFromMappack(id string) (models.TimeGoal, error)
	UpdateMappackTimeGoals(timegoals *[]models.TimeGoal) error
}

type mappackRepository struct {
	db *gorm.DB
}

func NewMappackRepository(db *gorm.DB) MappackRepository {
	return &mappackRepository{db: db}
}

func (t *mappackRepository) Create(mappack *models.Mappack) error {
	return t.db.Create(&mappack).Error

}
func (t *mappackRepository) GetById(id string) (models.Mappack, error) {
	mappack := models.Mappack{}
	err := t.db.Where("ID = ?", id).First(&mappack).Error
	return mappack, err
}

func (t *mappackRepository) GetAll() ([]models.Mappack, error) {
	mappacks := []models.Mappack{}
	err := t.db.Find(&mappacks).Error
	return mappacks, err
}

func (t *mappackRepository) CreateMappackTimeGoal(timegoal *models.TimeGoal) error {
	return t.db.Create(timegoal).Error
}
func (t *mappackRepository) GetAllMappackTimeGoals(mappackId string) ([]models.TimeGoal, error) {
	timegoals := []models.TimeGoal{}
	err := t.db.Where("mappack_id = ?", mappackId).Find(&timegoals).Error
	return timegoals, err
}

func (t *mappackRepository) RemoveTimeGoalFromMappack(id string) (models.TimeGoal, error) {
	timegoals := models.TimeGoal{}
	err := t.db.Where("id = ?", id).Delete(&timegoals).Error
	return timegoals, err
}

func (t *mappackRepository) UpdateMappackTimeGoals(timegoals *[]models.TimeGoal) error {
	for _, timegoal := range *timegoals {
		t.db.Save(&timegoal)
	}
	return nil
}
