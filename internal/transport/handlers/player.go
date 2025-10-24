package handlers

import (
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	PlayerService *services.PlayerService
}

func (t *PlayerHandler) Create(c *gin.Context) {
	player := models.Player{}

	err := c.ShouldBind(&player)
	if err != nil {
		fmt.Printf("Error occured while binding Player during creation: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	result := t.PlayerService.Create(&player)

	if result.Error != nil {
		fmt.Printf("Error occured while creating a Player: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Creation Succesful")
	}

}

func (t *PlayerHandler) GetAll(c *gin.Context) {
	players := []models.Player{}
	result := t.PlayerService.GetAll(&players)
	if result.Error != nil {
		fmt.Printf("Error occured while getting Players: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, players)
	}
}

func (t *PlayerHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	player := models.Player{}
	result := t.PlayerService.GetById(&player, id)
	if result.Error != nil {
		fmt.Printf("Error occured while getting a Player by id: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, player)
	}
}
