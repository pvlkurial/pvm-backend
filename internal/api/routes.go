package api

import (
	"example/pvm-backend/internal/database"
	"example/pvm-backend/internal/services"
	"example/pvm-backend/internal/transport/handlers"
	auth "example/pvm-backend/internal/transport/middleware"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	jwtConfig := auth.JWTConfig{
		SecretKey:     "",
		TokenDuration: 24 * time.Hour,
	}

	oauthConfig := handlers.OAuthConfig{
		ClientID:     os.Getenv("TRACKMANIA_CLIENT_ID"),
		ClientSecret: os.Getenv("TRACKMANIA_CLIENT_SECRET"),
		RedirectURI:  os.Getenv("TRACKMANIA_REDIRECT_URI"),
	}

	r.GET("/auth/login", handlers.HandleTrackmaniaLogin(oauthConfig))
	r.GET("/auth/callback", handlers.HandleTrackmaniaCallback(r.DB, oauthConfig, jwtConfig))

	r.Use(auth.RequireAuth(r.DB, jwtConfig))
	{
		r.GET("/api/profile", handlers.HandleGetProfile)
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

		r.POST("/mappacks/:mappack_id/tracks/:track_id/timegoals/:timegoal_id", trackHandler.CreateTimeGoalsForTrack)

		r.POST("/records", recordHandler.Create)
	}
	r.Run("localhost:8080")
}
