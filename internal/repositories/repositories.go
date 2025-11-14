package repositories

type Repositories struct {
	mappackRepository MappackRepository
	playerRepository  PlayerRepository
	recordRepository  RecordRepository
	trackRepository   TrackRepository
}

func NewRepositories(mappackRepository MappackRepository, playerRepository PlayerRepository,
	recordRepository RecordRepository, trackRepository TrackRepository) *Repositories {
	return &Repositories{mappackRepository: mappackRepository, playerRepository: playerRepository,
		recordRepository: recordRepository, trackRepository: trackRepository}
}
