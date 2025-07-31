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
