package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/models/dtos"
	"fmt"

	"gorm.io/gorm"
)

type RecordService struct {
	RecordRepository *database.RecordRepository
	PlayerRepository *database.PlayerRepository
	TrackRepository  *database.TrackRepository
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

func (t *RecordService) GetTrackWithRecords(track *dtos.TrackInMappackDto, mappackId string, trackId string) error {
	var trackInDb models.Track
	if err := t.TrackRepository.GetById(&trackInDb, trackId).Error; err != nil {
		return err
	}

	var records []models.Record
	if err := t.RecordRepository.GetByTrackId(&records, trackId).Error; err != nil {
		return err
	}

	var mappackTrack models.MappackTrack
	if err := t.TrackRepository.GetTrackInMappackInfo(&mappackTrack, mappackId, trackId).Error; err != nil {
		return err
	}

	var trackTimeGoals []models.TimeGoalMappackTrack
	if err := t.RecordRepository.GetTrackTimeGoalsTimes(mappackTrack.ID, &trackTimeGoals).Error; err != nil {
		return err
	}

	timeGoalDtos := make([]dtos.TrackTimeGoalDto, 0, len(trackTimeGoals))
	for _, ttg := range trackTimeGoals {
		timeGoalDtos = append(timeGoalDtos, dtos.TrackTimeGoalDto{
			Name: ttg.TimeGoal.Name,
			Time: ttg.Time,
		})
	}

	*track = dtos.TrackInMappackDto{
		ID:                       trackInDb.ID,
		MapID:                    trackInDb.MapID,
		MapUID:                   trackInDb.MapUID,
		Name:                     trackInDb.Name,
		Author:                   trackInDb.Author,
		Submitter:                trackInDb.Submitter,
		AuthorScore:              trackInDb.AuthorScore,
		GoldScore:                trackInDb.GoldScore,
		SilverScore:              trackInDb.SilverScore,
		BronzeScore:              trackInDb.BronzeScore,
		CollectionName:           trackInDb.CollectionName,
		Filename:                 trackInDb.Filename,
		MapType:                  trackInDb.MapType,
		MapStyle:                 trackInDb.MapStyle,
		IsPlayable:               trackInDb.IsPlayable,
		CreatedWithGamepadEditor: trackInDb.CreatedWithGamepadEditor,
		CreatedWithSimpleEditor:  trackInDb.CreatedWithSimpleEditor,
		Timestamp:                trackInDb.Timestamp,
		FileURL:                  trackInDb.FileURL,
		ThumbnailURL:             trackInDb.ThumbnailURL,
		Time:                     mappackTrack.ID,
		Tier:                     mappackTrack.Tier,
		UpdatedAt:                trackInDb.UpdatedAt,
		Records:                  records,
		TimeGoals:                timeGoalDtos,
	}

	return nil
}
