package models

import "time"

type Track struct {
	ID        string `gorm:"primary_key"`
	Name      string
	UpdatedAt time.Time
	MappackID string
	Mappack   Mappack
}
