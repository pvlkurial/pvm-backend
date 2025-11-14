package models

import "time"

type MappackTrack struct {
	MappackID string `gorm:"primaryKey" json:"mappack_id"`
	TrackID   string `gorm:"primaryKey" json:"track_id"`
	// time goals for track in mappack
	TimeGoalMappackTrack []TimeGoalMappackTrack `gorm:"foreignKey:MappackID,TrackID;references:MappackID,TrackID"`
	CreatedAt            time.Time
	Tier                 string `json:"tier"`
}
