package controllers

import (
	"example/pvm-backend/internal/models"
	"example/pvm-backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MappackController struct {
	mappackService services.MappackService
}

func (t *MappackController) Create(c *gin.Context) {
	mappack := models.Mappack{}

	err := c.ShouldBind(&mappack)
	if err != nil {
		fmt.Printf("Error occured while binding Mappack during creation: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	err = t.mappackService.Create(&mappack)

	if err != nil {
		fmt.Printf("Error occured while creating a Mappack: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Creation Succesful")
	}
}

func (t *MappackController) GetById(c *gin.Context) {
	id := c.Param("mappack_id")
	result, err := t.mappackService.GetById(id)
	if err != nil {
		fmt.Printf("Error occured while getting a Mappack by id: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (t *MappackController) GetAll(c *gin.Context) {
	result, err := t.mappackService.GetAll()
	if err != nil {
		fmt.Printf("Error occured while getting Mappacks: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (t *MappackController) CreateMappackTimeGoal(c *gin.Context) {
	mappackId := c.Param("mappack_id")
	timegoal := models.TimeGoal{}
	err := c.ShouldBind(&timegoal)
	if err != nil {
		fmt.Printf("Error occured while binding Mappack during creation: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	timegoal.MappackID = mappackId

	err = t.mappackService.CreateMappackTimeGoal(&timegoal)

	if err != nil {
		fmt.Printf("Error occured while creating a TimeGoal: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "Creation Successful", "data": timegoal})
	}
}

func (t *MappackController) GetAllMappackTimeGoals(c *gin.Context) {
	mappackId := c.Param("mappack_id")
	result, err := t.mappackService.GetAllMappackTimeGoals(mappackId)
	if err != nil {
		fmt.Printf("Error occured while getting TimeGoals of a mappack: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (t *MappackController) RemoveTimeGoalFromMappack(c *gin.Context) {
	id := c.Param("timegoal_id")
	result, err := t.mappackService.RemoveTimeGoalFromMappack(id)

	if err != nil {
		fmt.Printf("Error occured while removing a Timegoal from a mappack: %s", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "Removed %s timegoal from a mappack succesfully", result)
	}
}
