package workers

import (
	"example/pvm-backend/internal/clients"
	"example/pvm-backend/internal/services"
	"example/pvm-backend/internal/utils/constants"
	"fmt"
	"time"
)

type NadeoWorker interface {
	Start()
}

type nadeoWorker struct {
	recordService services.RecordService
	trackService  services.TrackService
	nadeoClient   clients.NadeoAPIClient
}

func NewNadeoWorker(recordService services.RecordService,
	trackService services.TrackService, nadeoClient clients.NadeoAPIClient) NadeoWorker {
	return &nadeoWorker{recordService: recordService, trackService: trackService, nadeoClient: nadeoClient}
}

func (n *nadeoWorker) Start() {

	go func() {
		ticker := time.NewTicker(constants.FetchIntervalInHours * time.Hour)
		defer ticker.Stop()
		for {
			for range ticker.C {
				dbTracks, err := n.trackService.GetAll()
				if err != nil {
					fmt.Printf("Error in nadeo worker while getting tracks: %v\n", err)
				}
				for _, track := range dbTracks {
					for i := 0; i < constants.TimesOfRecordsFetchPerTrack; i++ {
						records, err := n.nadeoClient.FetchRecordsOfTrack(track.ID,
							constants.RecordsPerRequest, i*constants.RecordsPerRequest)
						if err != nil {
							fmt.Printf("Error in nadeo worker while fetching records: %v\n", err)
						}
						err = n.recordService.SaveFetchedRecords(&records)
						if err != nil {
							fmt.Printf("Error in nadeo worker while saving records: %v\n", err)
						}
						time.Sleep(constants.FetchIntervalDelayInSeconds * time.Second)
					}
				}
			}
		}
	}()

}
