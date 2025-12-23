package models

import "time"

type MappackTrack struct {
	MappackID string `gorm:"primaryKey" json:"mappack_id"`
	TrackID   string `gorm:"primaryKey" json:"track_id"`
	Track     Track  `json:"track" gorm:"foreignKey:TrackID"`
	// time goals for track in mappack
	TimeGoalMappackTrack []TimeGoalMappackTrack `gorm:"foreignKey:MappackID,TrackID;references:MappackID,TrackID"`
	CreatedAt            time.Time
	TierName             *string     `json:"tier_name"`
	Tier                 MappackTier `gorm:"foreignKey:TierName,MappackID;references:Name,MappackID" json:"tier"`
	MapStyle             *string     `json:"mapStyle"`
}
