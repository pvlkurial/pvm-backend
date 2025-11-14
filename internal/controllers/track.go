package controllers

import (
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/services"
	"example/pvm-backend/internal/transport/handlers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrackController struct {
	trackService services.TrackService
	TokenManager *handlers.TokenManager
}

func NewTrackController(trackService services.TrackService) *TrackController {
	return &TrackController{trackService: trackService}
}

func (t *TrackController) Create(c *gin.Context) {
	trackTemp := models.Track{}

	err := c.ShouldBind(&trackTemp)
	track := t.TokenManager.FetchTrackInfo(trackTemp.ID)
	if err != nil {
		fmt.Printf("Error occured while binding Track during creation: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	fmt.Printf(track.MapType)

	err = t.trackService.Create(track)

	if err != nil {
		fmt.Printf("Error occured while creating a Track: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Creation Succesful, %s", track)
	}
}

func (t *TrackController) GetById(c *gin.Context) {
	id := c.Param("track_id")
	track, err := t.trackService.GetById(id)
	if err != nil {
		fmt.Printf("Error occured while getting a Track by id: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, track)
	}
}

func (t *TrackController) GetByMappackId(c *gin.Context) {
	id := c.Param("mappack_id")
	tracks, err := t.trackService.GetByMappackId(id)
	if err != nil {
		fmt.Printf("Error occured while getting a Tracks from a mappack by id: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, tracks)
	}
}

func (t *TrackController) AddTrackToMappack(c *gin.Context) {
	trackId := c.Param("track_id")
	mappackId := c.Param("mappack_id")

	err := t.trackService.AddTrackToMappack(trackId, mappackId)

	if err != nil {
		fmt.Printf("Error occured while creating a Track: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Added track to mappack succesfully")
	}
}

func (t *TrackController) RemoveTrackFromMappack(c *gin.Context) {
	trackId := c.Param("track_id")
	mappackId := c.Param("mappack_id")

	err := t.trackService.RemoveTrackFromMappack(trackId, mappackId)

	if err != nil {
		fmt.Printf("Error occured while removing a Track from mappack: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Removed track to mappack succesfully")
	}
}

func (t *TrackController) CreateTimeGoalsForTrack(c *gin.Context) {
	var timegoals []models.TimeGoalMappackTrack

	err := c.ShouldBind(&timegoals)
	if err != nil {
		fmt.Printf("Error occured while binding timegoals during creation/adding: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	err = t.trackService.CreateTimeGoalsForTrack(&timegoals)

	if err != nil {
		fmt.Printf("Error occured while creating a timegoal: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Creation Succesful")
	}
}

func (t *TrackController) GetTimeGoalsForTrack(c *gin.Context) {
	trackId := c.Param("track_id")
	mappackId := c.Param("mappack_id")

	timegoals, err := t.trackService.GetTimeGoalsForTrack(trackId, mappackId)

	if err != nil {
		fmt.Printf("Error occured while getting timegoals for track: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, timegoals)
	}
}

func (t *TrackController) UpdateTimeGoalsForTrack(c *gin.Context) {
	var timegoals []models.TimeGoalMappackTrack

	err := c.ShouldBind(&timegoals)
	if err != nil {
		fmt.Printf("Error occured while binding timegoals during update: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	err = t.trackService.UpdateTimeGoalsForTrack(&timegoals)

	if err != nil {
		fmt.Printf("Error occured while updating timegoals for track: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Update Succesful")
	}
}
