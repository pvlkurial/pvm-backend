package controllers

type Controllers struct {
	mappackController MappackController
	playerController  PlayerController
	recordController  RecordController
	trackController   TrackController
}

func NewControllers(mappackController MappackController, playerController PlayerController,
	recordController RecordController, trackController TrackController) *Controllers {
	return &Controllers{mappackController: mappackController, playerController: playerController,
		recordController: recordController, trackController: trackController}
}
