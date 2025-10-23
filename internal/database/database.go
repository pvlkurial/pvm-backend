package database

import (
	"example/pvm-backend/internal/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func ConnectDatabase() *gorm.DB {
	godotenv.Load(".env")
	dsn := os.Getenv("CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}
	return db
}

func MigrateDatabase(db *gorm.DB) {

	db.AutoMigrate(&models.TimeGoalMappackTrack{})
	db.AutoMigrate(&models.Mappack{})
	db.AutoMigrate(&models.MappackTrack{})
	db.AutoMigrate(&models.Track{})
	db.AutoMigrate(&models.Record{})
	db.AutoMigrate(&models.Player{})
	db.AutoMigrate(&models.TimeGoal{})
}
