package services

import "example/pvm-backend/internal/repositories"

type Services struct {
	MappackService MappackService
	PlayerService  PlayerService
	RecordService  RecordService
	TracksService  TrackService
}

func NewServices(repositories repositories.Repositories) *Services {
	mappackService := NewMappackService(repositories.MappackRepository)
	playerService := NewPlayerService(repositories.PlayerRepository)
	recordService := NewRecordService(repositories.RecordRepository,
		repositories.PlayerRepository, repositories.TrackRepository)
	trackService := NewTrackService(repositories.TrackRepository)

	return &Services{
		MappackService: mappackService,
		PlayerService:  playerService,
		RecordService:  recordService,
		TracksService:  trackService,
	}
}
