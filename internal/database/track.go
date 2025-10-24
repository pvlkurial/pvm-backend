package database

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type TrackRepository struct {
	DB *gorm.DB
}

func (t *TrackRepository) Create(track *models.Track) *gorm.DB {
	return t.DB.Create(&track)
}

func (t *TrackRepository) GetById(track *models.Track, id string) *gorm.DB {
	return t.DB.First(track).Where("ID = ?", id)
}

func (t *TrackRepository) GetByMappackId(tracks *[]models.Track, id string) *gorm.DB {
	return t.DB.Joins("JOIN mappack_tracks ON mappack_tracks.track_id = tracks.id AND mappack_tracks.mappack_id = ?", id).Find(tracks)
}

func (t *TrackRepository) AddTrackToMappack(mappackTrack *models.MappackTrack) *gorm.DB {
	return t.DB.Create(&mappackTrack)
}

func (t *TrackRepository) RemoveTrackFromMappack(trackId string, mappackId string) *gorm.DB {
	var mappackTrack models.MappackTrack
	res := t.DB.Select("id").Where("track_id = ?", trackId).Where("mappack_id = ?", mappackId).First(&mappackTrack)
	if res != nil {
		t.DB.Where("mappack_track_id = ?", mappackTrack.ID).Delete(&models.TimeGoalMappackTrack{})
	}
	return t.DB.Where("track_id = ?", trackId).Where("mappack_id = ?", mappackId).Delete(&models.MappackTrack{})
}

func (t *TrackRepository) CreateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) *gorm.DB {
	return t.DB.Create(timegoals)
}
func (t *TrackRepository) GetTimeGoalsForTrack(trackId string, mappackId string, timegoals *[]models.TimeGoalMappackTrack) *gorm.DB {
	var mappackTrack models.MappackTrack
	res := t.DB.Select("id").Where("track_id = ?", trackId).Where("mappack_id = ?", mappackId).First(&mappackTrack)
	if res != nil {
		return t.DB.Where("mappack_track_id = ?", mappackTrack.ID).Find(timegoals)
	}
	return res
}

func (t *TrackRepository) UpdateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) *gorm.DB {
	for _, timegoal := range *timegoals {
		t.DB.Save(&timegoal)
	}
	return nil
}
