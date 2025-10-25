package dtos

import (
	"example/pvm-backend/internal/models"

	"time"
)

type TrackInMappackDto struct {
	ID                       string             `json:"id"`
	MapID                    string             `json:"mapId"`
	MapUID                   string             `json:"mapUid"`
	Name                     string             `json:"name"`
	Author                   string             `json:"author"`
	Submitter                string             `json:"submitter"`
	AuthorScore              int                `json:"authorScore"`
	GoldScore                int                `json:"goldScore"`
	SilverScore              int                `json:"silverScore"`
	BronzeScore              int                `json:"bronzeScore"`
	CollectionName           string             `json:"collectionName"`
	Filename                 string             `json:"filename"`
	MapType                  string             `json:"mapType"`
	MapStyle                 string             `json:"mapStyle"`
	IsPlayable               bool               `json:"isPlayable"`
	CreatedWithGamepadEditor bool               `json:"createdWithGamepadEditor"`
	CreatedWithSimpleEditor  bool               `json:"createdWithSimpleEditor"`
	Timestamp                time.Time          `json:"timestamp"`
	FileURL                  string             `json:"fileUrl"`
	ThumbnailURL             string             `json:"thumbnailUrl"`
	Time                     int                `json:"time"`
	Tier                     string             `json:"tier"`
	UpdatedAt                time.Time          `json:"updatedAt"`
	Records                  []models.Record    `json:"records,omitempty"`
	TimeGoals                []TrackTimeGoalDto `json:"timegoals"`
}
