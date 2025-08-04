package models

import "time"

type MappackTrack struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	MappackID string `json:"mappack_id"`
	TrackID   string `json:"track_id"`
	// time goals for track in mappack
	TimeGoalMappackTrack []*TimeGoalMappackTrack
	CreatedAt            time.Time
}
