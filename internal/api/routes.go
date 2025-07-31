package api

import (
	"example/pvm-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Routes struct {
}

func (r Routes) InitRoutes() {
	router := gin.Default()
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Mappack{}, &models.Track{}, &models.Record{}, models.Player{})

	router.POST("/")
	router.Run("localhost:8080")
}
