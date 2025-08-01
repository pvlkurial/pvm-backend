package handlers

import (
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MappackHandler struct {
	MappackService *services.MappackService
}

func (t *MappackHandler) Create(c *gin.Context) {
	mappack := models.Mappack{}

	err := c.ShouldBind(&mappack)
	if err != nil {
		fmt.Printf("Error occured while binding Mappack during creation: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	result := t.MappackService.Create(&mappack)

	if result.Error != nil {
		fmt.Printf("Error occured while creating a Mappack: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Creation Succesful")
	}
}
