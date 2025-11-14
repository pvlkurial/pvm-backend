package api

import (
	"example/pvm-backend/internal/clients"
	"example/pvm-backend/internal/controllers"
	"example/pvm-backend/internal/repositories"
	"example/pvm-backend/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	*gin.Engine
	DB *gorm.DB
}

func (r *Routes) InitRoutes() {
	nadeoClient := clients.NewNadeoAPIClient()
	repositories := repositories.NewRepositories(r.DB)
	services := services.NewServices(*repositories, nadeoClient)
	controllers := controllers.NewControllers(*services, nadeoClient)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/tracks", controllers.TrackController.Create)
	r.GET("/tracks/:track_id", controllers.TrackController.GetById)
	r.DELETE("/tracks/:track_id")

	r.POST("/players", controllers.PlayerController.Create)
	r.GET("/players", controllers.PlayerController.GetAll)

	r.POST("/mappacks/:mappack_id/timegoals", controllers.MappackController.CreateMappackTimeGoal)
	r.GET("/mappacks/:mappack_id/timegoals", controllers.MappackController.GetAllMappackTimeGoals)
	r.DELETE("/mappacks/:mappack_id/timegoals/:timegoal_id", controllers.MappackController.RemoveTimeGoalFromMappack)

	r.POST("/mappacks", controllers.MappackController.Create)
	r.GET("/mappacks", controllers.MappackController.GetAll)
	r.GET("/mappacks/:mappack_id", controllers.MappackController.GetById)

	r.GET("/mappacks/:mappack_id/tracks", controllers.TrackController.GetByMappackId)
	r.POST("/mappacks/:mappack_id/tracks/:track_id", controllers.TrackController.AddTrackToMappack)
	r.DELETE("/mappacks/:mappack_id/tracks/:track_id", controllers.TrackController.RemoveTrackFromMappack)

	r.POST("/mappacks/:mappack_id/tracks/:track_id/timegoals", controllers.TrackController.CreateTimeGoalsForTrack)
	r.GET("/mappacks/:mappack_id/tracks/:track_id/timegoals", controllers.TrackController.GetTimeGoalsForTrack)
	r.PATCH("/mappacks/:mappack_id/tracks/:track_id/timegoals", controllers.TrackController.UpdateTimeGoalsForTrack)

	r.POST("/records", controllers.RecordController.Create)
	r.POST("/tracks/:track_id/records", controllers.RecordController.FetchNewTrackRecords)
	r.GET("/tracks/:track_id/records", controllers.RecordController.GetByTrackId)
	r.POST("/tracks/track_id/records/:player_id", controllers.RecordController.GetPlayersRecordsForTrack)

	r.GET("mappacks/:mappack_id/tracks/:track_id", controllers.RecordController.GetTrackWithRecords)

	r.Run("localhost:8080")
}
