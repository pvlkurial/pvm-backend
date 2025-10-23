package database

import (
	"example/pvm-backend/internal/models"
	"fmt"
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
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := "host=" + host + " user=" + user + " password=" + pw + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=Europe/Paris"
	fmt.Print(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
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
