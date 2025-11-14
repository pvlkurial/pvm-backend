package services

type Services struct {
	mappackService MappackService
	playerService  PlayerService
	recordService  RecordService
	tracksService  TrackService
}

func NewServices(mappackService MappackService, playerService PlayerService,
	recordService RecordService, tracksService TrackService) *Services {
	return &Services{mappackService: mappackService, playerService: playerService,
		recordService: recordService, tracksService: tracksService}
}
