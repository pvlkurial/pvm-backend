package services

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/models"
)

type TrackService interface {
	Create(track *models.Track) error
	GetById(id string) (models.Track, error)
	GetByMappackId(id string) ([]models.Track, error)
	AddTrackToMappack(trackId string, mappackId string) error
	RemoveTrackFromMappack(trackId string, mappackId string) error
	CreateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) error
	GetTimeGoalsForTrack(trackId string, mappackId string) ([]models.TimeGoalMappackTrack, error)
	UpdateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) error
	GetByUID(uid string) (models.Track, error)
}

type trackService struct {
	trackRepository database.TrackRepository
}

func NewTrackService(repo database.TrackRepository) TrackService {
	return &trackService{trackRepository: repo}
}

func (t *trackService) Create(track *models.Track) error {
	return t.trackRepository.Create(track)
}

func (t *trackService) GetById(id string) (models.Track, error) {
	return t.trackRepository.GetById(id)
}

func (t *trackService) GetByMappackId(id string) ([]models.Track, error) {
	return t.trackRepository.GetByMappackId(id)
}

func (t *trackService) AddTrackToMappack(trackId string, mappackId string) error {
	mappackTrack := models.MappackTrack{
		MappackID: mappackId,
		TrackID:   trackId,
	}
	return t.trackRepository.AddTrackToMappack(&mappackTrack)
}

func (t *trackService) RemoveTrackFromMappack(trackId string, mappackId string) error {
	return t.trackRepository.RemoveTrackFromMappack(trackId, mappackId)
}

func (t *trackService) CreateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) error {
	return t.trackRepository.CreateTimeGoalsForTrack(timegoals)
}

func (t *trackService) GetTimeGoalsForTrack(trackId string, mappackId string) ([]models.TimeGoalMappackTrack, error) {
	return t.trackRepository.GetTimeGoalsForTrack(trackId, mappackId)
}

func (t *trackService) UpdateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) error {
	return t.trackRepository.UpdateTimeGoalsForTrack(timegoals)
}

func (t *trackService) GetByUID(uid string) (models.Track, error) {
	return t.trackRepository.GetByUID(uid)
}
