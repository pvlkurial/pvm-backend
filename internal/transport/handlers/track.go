package handlers

import (
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrackHandler struct {
	TrackService *services.TrackService
}

func (t *TrackHandler) Create(c *gin.Context) {
	track := models.Track{}

	err := c.ShouldBind(&track)
	if err != nil {
		fmt.Printf("Error occured while binding Track during creation: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	result := t.TrackService.Create(&track)

	if result.Error != nil {
		fmt.Printf("Error occured while creating a Track: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Creation Succesful")
	}
}

func (t *TrackHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	track := models.Track{}
	result := t.TrackService.GetById(&track, id)
	if result.Error != nil {
		fmt.Printf("Error occured while getting a Track by id: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, track)
	}
}

func (t *TrackHandler) GetByMappackId(c *gin.Context) {
	id := c.Param("id")
	tracks := []models.Track{}
	result := t.TrackService.GetByMappackId(&tracks, id)
	if result.Error != nil {
		fmt.Printf("Error occured while getting a Tracks from a mappack by id: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, tracks)
	}
}

func (t *TrackHandler) AddTrackToMappack(c *gin.Context) {
	trackId := c.Param("track_id")
	mappackId := c.Param("mappack_id")

	result := t.TrackService.AddTrackToMappack(trackId, mappackId)

	if result.Error != nil {
		fmt.Printf("Error occured while creating a Track: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Added track to mappack succesfully")
	}
}

func (t *TrackHandler) RemoveTrackFromMappack(c *gin.Context) {
	trackId := c.Param("track_id")
	mappackId := c.Param("mappack_id")

	result := t.TrackService.RemoveTrackFromMappack(trackId, mappackId)

	if result.Error != nil {
		fmt.Printf("Error occured while removing a Track from mappack: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Removed track to mappack succesfully")
	}
}

func (t *TrackHandler) CreateTimeGoalsForTrack(c *gin.Context) {
	var timegoals []models.TimeGoalMappackTrack

	err := c.ShouldBind(&timegoals)
	if err != nil {
		fmt.Printf("Error occured while binding timegoals during creation/adding: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	result := t.TrackService.CreateTimeGoalsForTrack(&timegoals)

	if result.Error != nil {
		fmt.Printf("Error occured while creating a timegoal: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Creation Succesful")
	}
}
