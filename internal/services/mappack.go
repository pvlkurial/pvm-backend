package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type MappackService struct {
	MappackRepository *database.MappackRepository
}

func (t *MappackService) Create(mappack *models.Mappack) *gorm.DB {
	return t.MappackRepository.Create(mappack)
}

func (t *MappackService) GetById(mappack *models.Mappack, id string) *gorm.DB {
	return t.MappackRepository.GetById(mappack, id)
}
func (t *MappackService) GetAll(mappacks *[]models.Mappack) *gorm.DB {
	return t.MappackRepository.GetAll(mappacks)
}
func (t *MappackService) CreateMappackTimeGoal(timegoal *models.TimeGoal) *gorm.DB {
	return t.MappackRepository.CreateMappackTimeGoal(timegoal)
}
