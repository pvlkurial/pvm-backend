package models

type MappackTier struct {
	Name       string `gorm:"primaryKey" json:"name"`
	MappackID  string `gorm:"primaryKey" json:"mappack_id"`
	Multiplier int    `json:"multiplier"`
	Color      string `json:"color"`
}
