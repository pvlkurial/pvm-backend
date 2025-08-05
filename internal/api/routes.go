package api

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/services"
	"example/pvm-backend/internal/transport/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	*gin.Engine
	DB *gorm.DB
}

func (r *Routes) InitRoutes() {

	trackRepository := &database.TrackRepository{DB: r.DB}
	trackService := &services.TrackService{TrackRepository: trackRepository}
	trackHandler := &handlers.TrackHandler{TrackService: trackService}

	playerRepository := &database.PlayerRepository{DB: r.DB}
	playerService := &services.PlayerService{PlayerRepository: playerRepository}
	playerHandler := &handlers.PlayerHandler{PlayerService: playerService}

	mappackRepository := &database.MappackRepository{DB: r.DB}
	mappackService := &services.MappackService{MappackRepository: mappackRepository}
	mappackHandler := &handlers.MappackHandler{MappackService: mappackService}

	recordRepository := &database.RecordRepository{DB: r.DB}
	recordService := &services.RecordService{RecordRepository: recordRepository}
	recordHandler := &handlers.RecordHandler{RecordService: recordService}

	r.POST("/tracks", trackHandler.Create)
	r.GET("/tracks/:id", trackHandler.GetById)
	r.DELETE("/tracks/:id")

	r.POST("/players", playerHandler.Create)
	r.GET("/players", playerHandler.GetAll)

	r.POST("/mappacks", mappackHandler.Create)
	r.GET("/mappacks", mappackHandler.GetAll)
	r.GET("/mappacks/:id", mappackHandler.GetById)
	r.GET("/mappacks/:id/tracks", trackHandler.GetByMappackId)
	r.POST("/mappacks/:mappack_id/tracks/:track_id", trackHandler.AddTrackToMappack)
	r.DELETE("/mappacks/:mappack_id/tracks/:track_id", trackHandler.RemoveTrackFromMappack)

	r.POST("/records", recordHandler.Create)
	r.Run("localhost:8080")
}
