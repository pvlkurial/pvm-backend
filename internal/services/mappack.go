package services

import (
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/repositories"
)

type MappackService interface {
	Create(mappack *models.Mappack) error
	GetById(id string) (models.Mappack, error)
	GetAll() ([]models.Mappack, error)
	CreateMappackTimeGoal(timegoal *models.TimeGoal) error
	GetAllMappackTimeGoals(mappackId string) ([]models.TimeGoal, error)
	RemoveTimeGoalFromMappack(id string) (models.TimeGoal, error)
}

type mappackService struct {
	mappackRepository repositories.MappackRepository
}

func NewMappackService(repo repositories.MappackRepository) MappackService {
	return &mappackService{mappackRepository: repo}
}

func (t *mappackService) Create(mappack *models.Mappack) error {
	return t.mappackRepository.Create(mappack)
}

func (t *mappackService) GetById(id string) (models.Mappack, error) {
	return t.mappackRepository.GetById(id)
}
func (t *mappackService) GetAll() ([]models.Mappack, error) {
	return t.mappackRepository.GetAll()
}
func (t *mappackService) CreateMappackTimeGoal(timegoal *models.TimeGoal) error {
	return t.mappackRepository.CreateMappackTimeGoal(timegoal)
}
func (t *mappackService) GetAllMappackTimeGoals(mappackId string) ([]models.TimeGoal, error) {
	return t.mappackRepository.GetAllMappackTimeGoals(mappackId)
}
func (t *mappackService) RemoveTimeGoalFromMappack(id string) (models.TimeGoal, error) {
	return t.mappackRepository.RemoveTimeGoalFromMappack(id)
}
