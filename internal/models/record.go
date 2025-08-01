package models

import "time"

type Record struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	RecordTime int       `json:"record_time"`
	UpdatedAt  time.Time `json:"updated_at"`
	PlayerID   string    `type:"uuid"`
	TrackID    string
	Player     Player
	Track      Track
}
