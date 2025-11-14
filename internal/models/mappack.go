package models

import "time"

type Mappack struct {
	ID           string `gorm:"primaryKey" json:"id"`
	MappackTrack []*MappackTrack
	Name         string `json:"name"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsActive     bool `json:"isActive"`
	TimeGoals    []*TimeGoal
	MapStyle     MapStyle `gorm:"foreignKey:Name" json:"mapStyle"`
}
