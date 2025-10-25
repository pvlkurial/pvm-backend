package database

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type RecordRepository struct {
	DB *gorm.DB
}

func (t *RecordRepository) Create(record *models.Record) *gorm.DB {
	return t.DB.Save(&record)
}

func (t *RecordRepository) GetById(record *models.Record, id string) *gorm.DB {
	return t.DB.First(record).Where("ID = ?", id)
}

func (t *RecordRepository) GetByTrackId(records *[]models.Record, id string) *gorm.DB {
	subQuery := t.DB.Table("records").
		Select("player_id, MAX(updated_at) as max_time").
		Where("track_id = ?", id).
		Group("player_id")

	return t.DB.Joins("INNER JOIN (?) as latest ON records.player_id = latest.player_id AND records.updated_at = latest.max_time", subQuery).
		Where("records.track_id = ?", id).
		Find(records)
}

func (t *RecordRepository) GetPlayersRecordsForTrack(trackId string, playerId string, records *[]models.Record) *gorm.DB {
	return t.DB.Where("track_id = ?", trackId).Where("player_id = ?", playerId).Find(records)
}

func (t *RecordRepository) GetTrackTimeGoalsTimes(mappackTrackId int, trackTimeGoals *[]models.TimeGoalMappackTrack) *gorm.DB {
	return t.DB.
		Preload("TimeGoal").
		Where("mappack_track_id = ?", mappackTrackId).
		Find(trackTimeGoals)
}
