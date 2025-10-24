package api

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/services"
	"example/pvm-backend/internal/transport/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	*gin.Engine
	DB *gorm.DB
}

func (r *Routes) InitRoutes() {

	tokenManager := handlers.TokenManager{}
	tokenManager.AccessToken = ""

	trackRepository := &database.TrackRepository{DB: r.DB}
	trackService := &services.TrackService{TrackRepository: trackRepository}
	trackHandler := &handlers.TrackHandler{TrackService: trackService, TokenManager: &tokenManager}

	playerRepository := &database.PlayerRepository{DB: r.DB}
	playerService := &services.PlayerService{PlayerRepository: playerRepository}
	playerHandler := &handlers.PlayerHandler{PlayerService: playerService}

	mappackRepository := &database.MappackRepository{DB: r.DB}
	mappackService := &services.MappackService{MappackRepository: mappackRepository}
	mappackHandler := &handlers.MappackHandler{MappackService: mappackService}

	recordRepository := &database.RecordRepository{DB: r.DB}
	recordService := &services.RecordService{RecordRepository: recordRepository}
	recordHandler := &handlers.RecordHandler{RecordService: recordService}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/tracks", trackHandler.Create)
	r.GET("/tracks/:id", trackHandler.GetById)
	r.DELETE("/tracks/:id")

	r.POST("/players", playerHandler.Create)
	r.GET("/players", playerHandler.GetAll)

	r.POST("/mappacks/:mappack_id/timegoals", mappackHandler.CreateMappackTimeGoal)
	r.GET("/mappacks/:mappack_id/timegoals", mappackHandler.GetAllMappackTimeGoals)
	r.DELETE("/mappacks/:mappack_id/timegoals/:timegoal_id", mappackHandler.RemoveTimeGoalFromMappack)

	r.POST("/mappacks", mappackHandler.Create)
	r.GET("/mappacks", mappackHandler.GetAll)
	r.GET("/mappacks/:mappack_id", mappackHandler.GetById)

	r.GET("/mappacks/:mappack_id/tracks", trackHandler.GetByMappackId)
	r.POST("/mappacks/:mappack_id/tracks/:track_id", trackHandler.AddTrackToMappack)
	r.DELETE("/mappacks/:mappack_id/tracks/:track_id", trackHandler.RemoveTrackFromMappack)

	r.POST("/mappacks/:mappack_id/tracks/:track_id/timegoals", trackHandler.CreateTimeGoalsForTrack)
	r.GET("/mappacks/:mappack_id/tracks/:track_id/timegoals", trackHandler.GetTimeGoalsForTrack)
	r.PATCH("/mappacks/:mappack_id/tracks/:track_id/timegoals", trackHandler.UpdateTimeGoalsForTrack)

	r.POST("/records", recordHandler.Create)
	r.POST("/tracks/track_id/records", recordHandler.GetByTrackId)
	r.POST("/tracks/track_id/records/:player_id", recordHandler.GetPlayersRecordsForTrack)

	r.Run("localhost:8080")
}
