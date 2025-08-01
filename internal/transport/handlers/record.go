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
