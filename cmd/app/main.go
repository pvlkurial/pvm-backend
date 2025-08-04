package main

import (
	"example/pvm-backend/internal/api"
	"example/pvm-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// db.Migrator().DropTable(&models.Mappack{}, &models.Track{}, &models.Record{}, models.Player{},
	// 	models.MappackTrack{}, models.TimeGoal{}, models.TimeGoalMappackTrack{})
	db.AutoMigrate(&models.TimeGoalMappackTrack{})
	db.AutoMigrate(&models.Mappack{})

	db.AutoMigrate(&models.MappackTrack{})
	db.AutoMigrate(&models.Track{})
	db.AutoMigrate(&models.Record{})
	db.AutoMigrate(&models.Player{})
	db.AutoMigrate(&models.TimeGoal{})
	r := api.Routes{router, db}
	r.InitRoutes()
}
