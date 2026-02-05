package services

import (
	"example/pvm-backend/internal/clients"
	"example/pvm-backend/internal/repositories"
)

type Services struct {
	MappackService     MappackService
	PlayerService      PlayerService
	RecordService      RecordService
	TracksService      TrackService
	AchievementService AchievementService
}

func NewServices(repositories repositories.Repositories, client *clients.NadeoAPIClient, tmClient clients.TrackmaniaAPIClient) *Services {
	mappackService := NewMappackService(repositories.MappackRepository, repositories.PlayerRepository, tmClient)
	playerService := NewPlayerService(repositories.PlayerRepository)
	trackService := NewTrackService(repositories.TrackRepository, client)
	achievementService := NewAchievementService(repositories.AchievementRepository, repositories.TrackRepository)
	recordService := NewRecordService(repositories.RecordRepository,
		repositories.PlayerRepository, repositories.TrackRepository, tmClient, achievementService)

	return &Services{
		MappackService:     mappackService,
		PlayerService:      playerService,
		RecordService:      recordService,
		TracksService:      trackService,
		AchievementService: *achievementService,
	}
}
