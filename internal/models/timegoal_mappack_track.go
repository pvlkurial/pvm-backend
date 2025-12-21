package models

import "time"

// "Instance" of Time goal type for a track in a mappack
type TimeGoalMappackTrack struct {
	TimeGoalID int    `gorm:"primaryKey"`
	MappackID  string `gorm:"primaryKey"`
	TrackID    string `gorm:"primaryKey"`
	Time       int    `json:"time"`
	Multiplier int    `json:"multiplier"`
	UpdatedAt  time.Time
	TimeGoal   TimeGoal `gorm:"foreignKey:TimeGoalID"`
}
