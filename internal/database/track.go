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
