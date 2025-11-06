package controllers

import (
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/models/dtos"
	"example/pvm-backend/internal/services"
	"example/pvm-backend/internal/transport/handlers"
	"time"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordController struct {
	RecordService *services.RecordService
	TokenManager  *handlers.TokenManager
	TrackService  *services.TrackService
}

func (t *RecordController) Create(c *gin.Context) {
	record := models.Record{}

	err := c.ShouldBind(&record)
	if err != nil {
		fmt.Printf("Error occured while binding Record during creation: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	result := t.RecordService.Create(&record)

	if result.Error != nil {
		fmt.Printf("Error occured while creating a Record: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Creation Succesful")
	}

}

func (t *RecordController) GetById(c *gin.Context) {
	id := c.Param("id")
	record := models.Record{}
	result := t.RecordService.GetById(&record, id)
	if result.Error != nil {
		fmt.Printf("Error occured while getting a Record by id: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, record)
	}
}

func (t *RecordController) GetByTrackId(c *gin.Context) {
	trackId := c.Param("track_id")
	records := []models.Record{}
	track := models.Track{}
	trackResult := t.TrackService.GetById(&track, trackId)
	if trackResult.Error != nil {
		fmt.Printf("Error occured while getting Track by id: %s", trackResult.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	result := t.RecordService.GetByTrackId(&records, track.MapID)
	if result.Error != nil {
		fmt.Printf("Error occured while getting Records by Track id: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, records)
	}
}

func (t *RecordController) GetPlayersRecordsForTrack(c *gin.Context) {
	trackId := c.Param("track_id")
	playerId := c.Param("player_id")
	records := []models.Record{}
	result := t.RecordService.GetPlayersRecordsForTrack(trackId, playerId, &records)
	if result.Error != nil {
		fmt.Printf("Error occured while getting Player's Records for Track: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, records)
	}
}

func (t *RecordController) FetchNewTrackRecords(c *gin.Context) {
	trackId := c.Param("track_id")
	track := models.Track{}
	trackResult := t.TrackService.GetById(&track, trackId)
	if trackResult.Error != nil {
		fmt.Printf("Error occurred while fetching Track by ID: %s\n", trackResult.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	recordList := t.TokenManager.FetchRecordsOfTrack(track.MapUID)

	if recordList == nil {
		fmt.Println("Failed to fetch records")
		c.String(http.StatusInternalServerError, "Failed to fetch records")
		return
	}

	if len(*recordList) == 0 {
		c.String(http.StatusOK, "No records found")
		return
	}

	for i := range *recordList {
		(*recordList)[i].TrackID = track.ID
		(*recordList)[i].ID = fmt.Sprintf("%s_%s", track.ID, (*recordList)[i].PlayerID)
		(*recordList)[i].UpdatedAt = time.Now()
	}

	result := t.RecordService.SaveFetchedRecords(recordList)

	if result == nil {
		c.String(http.StatusOK, "No records to save")
		return
	}

	if result.Error != nil {
		fmt.Printf("Error occurred while creating a Record: %s\n", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Records saved successfully", "count": len(*recordList)})
}
func (t *RecordController) GetTrackWithRecords(c *gin.Context) {
	trackId := c.Param("track_id")
	mappack_id := c.Param("mappack_id")
	var track dtos.TrackInMappackDto
	t.RecordService.GetTrackWithRecords(&track, mappack_id, trackId)
	c.JSON(http.StatusOK, track)
}
