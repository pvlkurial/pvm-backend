package services

import (
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/models/dtos"
	"example/pvm-backend/internal/repositories"
	"fmt"
	"time"
)

type RecordService interface {
	Create(record *models.Record) error
	GetById(id string) (models.Record, error)
	GetByTrackId(id string) ([]models.Record, error)
	GetPlayersRecordsForTrack(trackId string, playerId string) ([]models.Record, error)
	SaveFetchedRecords(records *[]models.Record) error
	GetTrackWithRecords(mappackId string, trackId string) (dtos.TrackInMappackDto, error)
}

type recordService struct {
	recordRepository repositories.RecordRepository
	playerRepository repositories.PlayerRepository
	trackRepository  repositories.TrackRepository
}

func NewRecordService(recordRepo repositories.RecordRepository,
	playerRepo repositories.PlayerRepository,
	trackRepo repositories.TrackRepository) RecordService {
	return &recordService{recordRepository: recordRepo, playerRepository: playerRepo, trackRepository: trackRepo}
}

func (t *recordService) Create(record *models.Record) error {
	return t.recordRepository.Create(record)
}

func (t *recordService) GetById(id string) (models.Record, error) {
	return t.recordRepository.GetById(id)
}

func (t *recordService) GetByTrackId(id string) ([]models.Record, error) {
	return t.recordRepository.GetByTrackId(id)
}

func (t *recordService) GetPlayersRecordsForTrack(trackId string, playerId string) ([]models.Record, error) {
	return t.recordRepository.GetPlayersRecordsForTrack(trackId, playerId)
}

func (t *recordService) SaveFetchedRecords(records *[]models.Record) error {
	if records == nil || len(*records) == 0 {
		return nil
	}

	for _, record := range *records {
		_, err := t.playerRepository.GetById(record.PlayerID)
		if err != nil {
			fmt.Printf("Player %s not found.\n", record.PlayerID)
			err = t.playerRepository.Create(&models.Player{ID: record.PlayerID})
			if err != nil {
				fmt.Printf("Error creating player %s: %v\n", record.PlayerID, err)
				continue
			}
			fmt.Printf("Player %s created.\n", record.PlayerID)
		}

		err = t.recordRepository.Create(&record)
		if err != nil {
			fmt.Printf("Error creating record: %v\n", err)
			return err
		}
	}
	return nil
}

func (t *recordService) GetTrackWithRecords(mappackId string, trackId string) (dtos.TrackInMappackDto, error) {
	emptyTrack := dtos.TrackInMappackDto{}

	trackInDb, err := t.trackRepository.GetById(trackId)
	if err != nil {
		return emptyTrack, err
	}

	records, err := t.recordRepository.GetByTrackId(trackId)
	if err != nil {
		return emptyTrack, err
	}

	mappackTrack, err := t.trackRepository.GetTrackInMappackInfo(mappackId, trackId)
	if err != nil {
		return emptyTrack, err
	}

	trackTimeGoals, err := t.recordRepository.GetTrackTimeGoalsTimes(mappackId, trackId)
	if err != nil {
		return emptyTrack, err
	}

	timeGoalDtos := make([]dtos.TrackTimeGoalDto, 0, len(trackTimeGoals))
	for _, ttg := range trackTimeGoals {
		timeGoalDtos = append(timeGoalDtos, dtos.TrackTimeGoalDto{
			Name: ttg.TimeGoal.Name,
			Time: ttg.Time,
		})
	}

	track := dtos.TrackInMappackDto{
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
		Time:                     int(time.Now().Unix()),
		Tier:                     mappackTrack.Tier,
		UpdatedAt:                trackInDb.UpdatedAt,
		Records:                  records,
		TimeGoals:                timeGoalDtos,
	}
	return track, nil
}
