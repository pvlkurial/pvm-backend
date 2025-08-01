package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type RecordService struct {
	RecordRepository *database.RecordRepository
}

func (t *RecordService) Create(record *models.Record) *gorm.DB {
	return t.RecordRepository.Create(record)
}
