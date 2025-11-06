package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"
)

type MappackService struct {
	MappackRepository *database.MappackRepository
}

func (t *MappackService) Create(mappack *models.Mappack) error {

	return t.MappackRepository.Create(mappack)
}

func (t *MappackService) GetById(id string) (models.Mappack, error) {
	return t.MappackRepository.GetById(id)
}
func (t *MappackService) GetAll() ([]models.Mappack, error) {
	return t.MappackRepository.GetAll()
}
func (t *MappackService) CreateMappackTimeGoal(timegoal *models.TimeGoal) error {
	return t.MappackRepository.CreateMappackTimeGoal(timegoal)
}
func (t *MappackService) GetAllMappackTimeGoals(mappackId string) ([]models.TimeGoal, error) {
	return t.MappackRepository.GetAllMappackTimeGoals(mappackId)
}
func (t *MappackService) RemoveTimeGoalFromMappack(id string) (models.TimeGoal, error) {
	return t.MappackRepository.RemoveTimeGoalFromMappack(id)
}
