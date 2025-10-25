package models

import "time"

// "Instance" of Time goal type for a track in a mappack
type TimeGoalMappackTrack struct {
	ID             int    `gorm:"primaryKey"`
	TimeGoalID     int    `json:"time_goal_id"`
	MappackTrackID int    `json:"mappack_track_id"`
	Time           int    `json:"time"`
	Tier           string `json:"tier"`
	UpdatedAt      time.Time
}
