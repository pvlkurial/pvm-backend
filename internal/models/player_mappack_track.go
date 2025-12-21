package models

type PlayerMappackTrack struct {
	PlayerID         string               `gorm:"primaryKey" json:"playerId"`
	MappackTrackID   string               `gorm:"primaryKey" json:"mappackTrackId"`
	Score            int                  `json:"score"`
	AchievedTimeGoal TimeGoalMappackTrack `gorm:"foreignKey:TimeGoalID,MappackID,TrackID;"`
}
