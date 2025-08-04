package models

import "time"

// Time Goal Type: Alien Time, Bronze Time, etc...
type TimeGoal struct {
	ID                   int    `gorm:"primary_key"`
	Name                 string `json:"name"`
	MappackID            string `json:"mappack_id"`
	TimeGoalMappackTrack []*TimeGoalMappackTrack
	UpdatedAt            time.Time
}
