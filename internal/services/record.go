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

func (t *RecordService) GetById(record *models.Record, id string) *gorm.DB {
	return t.RecordRepository.GetById(record, id)
}

func (t *RecordService) GetByTrackId(records *[]models.Record, id string) *gorm.DB {
	return t.RecordRepository.GetByTrackId(records, id)
}

func (t *RecordService) GetPlayersRecordsForTrack(trackId string, playerId string, records *[]models.Record) *gorm.DB {
	return t.RecordRepository.GetPlayersRecordsForTrack(trackId, playerId, records)
}
