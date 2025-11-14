package controllers

import (
	"example/pvm-backend/internal/clients"
	"example/pvm-backend/internal/services"
)

type Controllers struct {
	MappackController MappackController
	PlayerController  PlayerController
	RecordController  RecordController
	TrackController   TrackController
}

func NewControllers(services services.Services, client *clients.NadeoAPIClient) *Controllers {
	mappackController := NewMappackController(services.MappackService)
	playerController := NewPlayerController(services.PlayerService)
	recordController := NewRecordController(services.RecordService, services.TracksService, client)
	trackController := NewTrackController(services.TracksService)

	return &Controllers{MappackController: *mappackController, PlayerController: *playerController,
		RecordController: *recordController, TrackController: *trackController}
}
