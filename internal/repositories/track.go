package repositories

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type TrackRepository interface {
	Create(track *models.Track) error
	GetById(id string) (models.Track, error)
	GetAll() ([]models.Track, error)
	GetByMappackId(id string) ([]models.Track, error)
	AddTrackToMappack(mappackTrack *models.MappackTrack) error
	RemoveTrackFromMappack(trackId string, mappackId string) error
	CreateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) error
	GetTimeGoalsForTrack(trackId string, mappackId string) ([]models.TimeGoalMappackTrack, error)
	UpdateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) error
	GetTrackInMappackInfo(mappackId string, trackId string) (models.MappackTrack, error)
	GetByUID(uid string) (models.Track, error)
	SavePlayerMappackTrack(playerMappackTrack *models.PlayerMappackTrack) error
}

type trackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB) TrackRepository {
	return &trackRepository{db: db}
}

func (t *trackRepository) Create(track *models.Track) error {
	return t.db.Save(track).Error
}

func (t *trackRepository) GetById(id string) (models.Track, error) {
	track := models.Track{}
	err := t.db.Where("ID = ?", id).First(&track).Error
	return track, err
}
func (t *trackRepository) GetAll() ([]models.Track, error) {
	var tracks []models.Track
	err := t.db.Find(&tracks).Error
	return tracks, err
}

func (t *trackRepository) GetByMappackId(id string) ([]models.Track, error) {
	tracks := []models.Track{}
	err := t.db.Joins("JOIN mappack_tracks ON mappack_tracks.track_id = tracks.id AND mappack_tracks.mappack_id = ?", id).Find(&tracks).Error
	return tracks, err
}

func (t *trackRepository) AddTrackToMappack(mappackTrack *models.MappackTrack) error {
	return t.db.Create(&mappackTrack).Error
}

func (t *trackRepository) RemoveTrackFromMappack(trackId string, mappackId string) error {
	err := t.db.Where("track_id = ? AND mappack_id = ?", trackId, mappackId).Delete(&models.TimeGoalMappackTrack{}).Error
	if err != nil {
		return err
	}
	return t.db.Where("track_id = ? AND mappack_id = ?", trackId, mappackId).Delete(&models.MappackTrack{}).Error
}

func (t *trackRepository) CreateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) error {
	return t.db.Create(timegoals).Error
}

func (t *trackRepository) GetTimeGoalsForTrack(trackId string, mappackId string) ([]models.TimeGoalMappackTrack, error) {
	var timegoals []models.TimeGoalMappackTrack
	err := t.db.Where("track_id = ? AND mappack_id = ?", trackId, mappackId).Find(&timegoals).Error
	return timegoals, err
}

func (t *trackRepository) UpdateTimeGoalsForTrack(timegoals *[]models.TimeGoalMappackTrack) error {
	return t.db.Save(timegoals).Error
}
func (t *trackRepository) GetTrackInMappackInfo(mappackId string, trackId string) (models.MappackTrack, error) {
	mappackTrack := models.MappackTrack{}
	err := t.db.Where("mappack_id = ? AND track_id = ?", mappackId, trackId).First(&mappackTrack).Error
	return mappackTrack, err
}

func (t *trackRepository) GetByUID(uid string) (models.Track, error) {
	track := models.Track{}
	err := t.db.Where("map_uid = ?", uid).First(&track).Error
	return track, err
}

func (t *trackRepository) SavePlayerMappackTrack(playerMappackTrack *models.PlayerMappackTrack) error {
	return t.db.Save(playerMappackTrack).Error
}
