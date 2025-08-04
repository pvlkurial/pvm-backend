package models

import "time"

type Track struct {
	ID           string `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	UpdatedAt    time.Time
	MappackTrack []*MappackTrack
	WorldRecord  int `json:"world_record"`
}
