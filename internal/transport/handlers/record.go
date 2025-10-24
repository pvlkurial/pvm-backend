package handlers

import (
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordHandler struct {
	RecordService *services.RecordService
}

func (t *RecordHandler) Create(c *gin.Context) {
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

func (t *RecordHandler) GetById(c *gin.Context) {
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

func (t *RecordHandler) GetByTrackId(c *gin.Context) {
	id := c.Param("track_id")
	records := []models.Record{}
	result := t.RecordService.GetByTrackId(&records, id)
	if result.Error != nil {
		fmt.Printf("Error occured while getting Records by Track id: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, records)
	}
}

func (t *RecordHandler) GetPlayersRecordsForTrack(c *gin.Context) {
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
