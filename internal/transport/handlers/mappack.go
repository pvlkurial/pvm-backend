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

func (t *MappackHandler) GetById(c *gin.Context) {
	id := c.Param("mappack_id")
	mappack := models.Mappack{}
	result := t.MappackService.GetById(&mappack, id)
	if result.Error != nil {
		fmt.Printf("Error occured while getting a Mappack by id: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, mappack)
	}
}

func (t *MappackHandler) GetAll(c *gin.Context) {
	mappacks := []models.Mappack{}
	result := t.MappackService.GetAll(&mappacks)
	if result.Error != nil {
		fmt.Printf("Error occured while getting Mappacks: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, mappacks)
	}
}

func (t *MappackHandler) CreateMappackTimeGoal(c *gin.Context) {
	mappackId := c.Param("mappack_id")
	timegoal := models.TimeGoal{}
	err := c.ShouldBind(&timegoal)
	if err != nil {
		fmt.Printf("Error occured while binding Mappack during creation: %s", err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	timegoal.MappackID = mappackId

	result := t.MappackService.CreateMappackTimeGoal(&timegoal)

	if result.Error != nil {
		fmt.Printf("Error occured while creating a TimeGoal: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "Creation Successful", "data": timegoal})
	}
}

func (t *MappackHandler) GetAllMappackTimeGoals(c *gin.Context) {
	mappackId := c.Param("mappack_id")
	timegoals := []models.TimeGoal{}
	result := t.MappackService.GetAllMappackTimeGoals(mappackId, &timegoals)
	if result.Error != nil {
		fmt.Printf("Error occured while getting TimeGoals of a mappack: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, timegoals)
	}
}

func (t *MappackHandler) RemoveTimeGoalFromMappack(c *gin.Context) {
	id := c.Param("timegoal_id")

	result := t.MappackService.RemoveTimeGoalFromMappack(id)

	if result.Error != nil {
		fmt.Printf("Error occured while removing a Timegoal from a mappack: %s", result.Error)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Removed timegoal from a mappack succesfully")
	}
}
