package repositories

import "gorm.io/gorm"

type Repositories struct {
	MappackRepository MappackRepository
	PlayerRepository  PlayerRepository
	RecordRepository  RecordRepository
	TrackRepository   TrackRepository
}

func NewRepositories(db *gorm.DB) *Repositories {

	mappackRepository := NewMappackRepository(db)
	playerRepository := NewPlayerRepository(db)
	recordRepository := NewRecordRepository(db)
	trackRepository := NewTrackRepository(db)

	return &Repositories{MappackRepository: mappackRepository, PlayerRepository: playerRepository,
		RecordRepository: recordRepository, TrackRepository: trackRepository}
}
