package repositories

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type RecordRepository interface {
	Create(record *models.Record) error
	GetById(id string) (models.Record, error)
	GetByTrackId(id string) ([]models.Record, error)
	GetPlayersRecordsForTrack(trackId string, playerId string) ([]models.Record, error)
	GetTrackTimeGoalsTimes(mappackTrackId int) ([]models.TimeGoalMappackTrack, error)
}

type recordRepository struct {
	DB *gorm.DB
}

func NewRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{DB: db}
}

func (t *recordRepository) Create(record *models.Record) error {
	return t.DB.Save(&record).Error
}

func (t *recordRepository) GetById(id string) (models.Record, error) {
	record := models.Record{}
	err := t.DB.First(&record).Where("ID = ?", id).Error
	return record, err
}

func (t *recordRepository) GetByTrackId(id string) ([]models.Record, error) {
	records := []models.Record{}
	subQuery := t.DB.Table("records").
		Select("player_id, MAX(updated_at) as max_time").
		Where("track_id = ?", id).
		Group("player_id")

	err := t.DB.Joins("INNER JOIN (?) as latest ON records.player_id = latest.player_id AND records.updated_at = latest.max_time", subQuery).
		Where("records.track_id = ?", id).
		Find(&records).Error

	return records, err
}

func (t *recordRepository) GetPlayersRecordsForTrack(trackId string, playerId string) ([]models.Record, error) {
	records := []models.Record{}
	err := t.DB.Where("track_id = ?", trackId).Where("player_id = ?", playerId).Find(&records).Error
	return records, err
}

func (t *recordRepository) GetTrackTimeGoalsTimes(mappackTrackId int) ([]models.TimeGoalMappackTrack, error) {
	trackTimeGoals := []models.TimeGoalMappackTrack{}
	err := t.DB.
		Preload("TimeGoal").
		Where("mappack_track_id = ?", mappackTrackId).
		Find(&trackTimeGoals).Error
	return trackTimeGoals, err
}
