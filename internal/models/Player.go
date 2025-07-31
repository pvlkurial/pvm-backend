package models

type Player struct {
	ID     string `gorm:"primary_key;type:uuid"`
	Name   string `json:"name"`
	Record []Record
}
