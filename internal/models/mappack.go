package models

import "time"

type Mappack struct {
	ID        string `gorm:"primary_key"`
	Track     []Track
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool
}
