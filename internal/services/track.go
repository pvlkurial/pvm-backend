package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type TrackService struct {
	TrackRepository *database.TrackRepository
}

func (t *TrackService) Create(track *models.Track) *gorm.DB {
	return t.TrackRepository.Create(track)
}

func (t *TrackService) GetById(track *models.Track, id string) *gorm.DB {
	return t.TrackRepository.GetById(track, id)
}

func (t *TrackService) GetByMappackId(tracks *[]models.Track, id string) *gorm.DB {
	return t.TrackRepository.GetByMappackId(tracks, id)
}

func (t *TrackService) AddTrackToMappack(trackId string, mappackId string) *gorm.DB {
	mappackTrack := models.MappackTrack{
		MappackID: mappackId,
		TrackID:   trackId,
	}
	return t.TrackRepository.AddTrackToMappack(&mappackTrack)
}

func (t *TrackService) RemoveTrackFromMappack(trackId string, mappackId string) *gorm.DB {
	return t.TrackRepository.RemoveTrackFromMappack(trackId, mappackId)
}

func (t *TrackService) CreateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) *gorm.DB {
	return t.TrackRepository.CreateTimeGoalsForTrack(timegoals)
}

func (t *TrackService) GetTimeGoalsForTrack(trackId string, mappackId string, timegoals *[]models.TimeGoalMappackTrack) *gorm.DB {
	return t.TrackRepository.GetTimeGoalsForTrack(trackId, mappackId, timegoals)
}

func (t *TrackService) UpdateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) *gorm.DB {
	return t.TrackRepository.UpdateTimeGoalsForTrack(timegoals)
}
