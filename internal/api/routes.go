package api

import (
	"example/pvm-backend/internal/clients"
	"example/pvm-backend/internal/controllers"
	"example/pvm-backend/internal/repositories"
	"example/pvm-backend/internal/services"
	"example/pvm-backend/internal/workers"

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
	trackmaniaClient := clients.NewTrackmaniaAPIClient()
	repositories := repositories.NewRepositories(r.DB)
	services := services.NewServices(*repositories, nadeoClient, *trackmaniaClient)
	controllers := controllers.NewControllers(*services, nadeoClient)

	workers := workers.NewWorkers(*services, *nadeoClient)
	workers.NadeoWorker.Start()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
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
	r.PUT("/mappacks/:mappack_id/timegoals", controllers.MappackController.UpdateMappackTimeGoals)
	//r.DELETE("/mappacks/:mappack_id/timegoals/:timegoal_id", controllers.MappackController.RemoveTimeGoalFromMappack)

	r.POST("/mappacks", controllers.MappackController.Create)
	r.PUT("/mappacks", controllers.MappackController.Update)
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

	r.GET("/mappacks/:mappack_id/leaderboard", controllers.AchievementController.GetMappackLeaderboard)
	r.GET("/mappacks/:mappack_id/players/:player_id/achievements", controllers.AchievementController.GetPlayerAchievements)
	r.GET("/mappacks/:mappack_id/players/:player_id/rank", controllers.AchievementController.GetPlayerRank)
	r.GET("/mappacks/:mappack_id/players/:player_id/leaderboard-entry", controllers.AchievementController.GetPlayerLeaderboardEntry)

	r.POST("/mappacks/:mappack_id/recalculate-achievements", controllers.AchievementController.RecalculateMappackAchievements)
	r.DELETE("/mappacks/:mappack_id/timegoals/:id", controllers.MappackController.DeleteTimeGoal)
	r.DELETE("/mappacks/:mappack_id/tiers/:id", controllers.MappackController.DeleteTier)
	r.DELETE("/mappacks/:mappack_id/ranks/:id", controllers.MappackController.DeleteRank)

	r.Run("localhost:8080")
}
