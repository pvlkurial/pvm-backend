package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type RecordService struct {
	RecordRepository *database.RecordRepository
	PlayerRepository *database.PlayerRepository
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

func (t *RecordService) SaveFetchedRecords(records *[]models.Record) *gorm.DB {
	if records == nil || len(*records) == 0 {
		return &gorm.DB{}
	}

	var result *gorm.DB
	for _, record := range *records {
		var player models.Player
		playerResult := t.RecordRepository.DB.First(&player, "id = ?", record.PlayerID)
		if playerResult.Error == gorm.ErrRecordNotFound {
			fmt.Printf("Player %s not found.\n", record.PlayerID)
			playerRes := t.PlayerRepository.Create(&models.Player{ID: record.PlayerID})
			if playerRes.Error != nil {
				fmt.Printf("Error creating player %s: %v\n", record.PlayerID, playerRes.Error)
				continue
			}
			fmt.Printf("Player %s created.\n", record.PlayerID)
		}

		result = t.RecordRepository.Create(&record)
		if result.Error != nil {
			fmt.Printf("Error creating record: %v\n", result.Error)
			return result
		}
	}
	return result
}
