package database

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type RecordRepository struct {
	DB *gorm.DB
}

func (t *RecordRepository) Create(record *models.Record) *gorm.DB {
	return t.DB.Create(&record)
}
