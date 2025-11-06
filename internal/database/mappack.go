package database

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type MappackRepository struct {
	DB *gorm.DB
}

func (t *MappackRepository) Create(mappack *models.Mappack) error {
	return t.DB.Create(&mappack).Error

}
func (t *MappackRepository) GetById(id string) (models.Mappack, error) {
	mappack := models.Mappack{}
	err := t.DB.Where("ID = ?", id).First(&mappack).Error
	return mappack, err
}

func (t *MappackRepository) GetAll() ([]models.Mappack, error) {
	mappacks := []models.Mappack{}
	err := t.DB.Find(&mappacks).Error
	return mappacks, err
}

func (t *MappackRepository) CreateMappackTimeGoal(timegoal *models.TimeGoal) error {
	return t.DB.Create(timegoal).Error
}
func (t *MappackRepository) GetAllMappackTimeGoals(mappackId string) ([]models.TimeGoal, error) {
	timegoals := []models.TimeGoal{}
	err := t.DB.Where("mappack_id = ?", mappackId).Find(&timegoals).Error
	return timegoals, err
}

func (t *MappackRepository) RemoveTimeGoalFromMappack(id string) (models.TimeGoal, error) {
	timegoals := models.TimeGoal{}
	err := t.DB.Where("id = ?", id).Delete(&timegoals).Error
	return timegoals, err
}

func (t *MappackRepository) UpdateMappackTimeGoals(timegoals *[]models.TimeGoal) error {
	for _, timegoal := range *timegoals {
		t.DB.Save(&timegoal)
	}
	return nil
}
