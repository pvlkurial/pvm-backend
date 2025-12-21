package seeds

import (
	"example/pvm-backend/internal/models"

	"gorm.io/gorm"
)

type TrackSeeder struct {
	DB *gorm.DB
}

func (m *TrackSeeder) seedTracks() error {
	tracks := []models.Track{
		{
			Author:                   "d2372a08-a8a1-46cb-97fb-23a161d85ad0",
			AuthorScore:              29089,
			BronzeScore:              44000,
			CollectionName:           "Stadium",
			CreatedWithGamepadEditor: false,
			CreatedWithSimpleEditor:  false,
			Filename:                 "Spring 2024 - 01.Map.Gbx",
			GoldScore:                31000,
			IsPlayable:               true,
			MapID:                    "d2b8a048-209d-4cfa-b5a4-bc3e3cab3566",
			MapStyle:                 "",
			MapType:                  `TrackMania\\TM_Race`,
			MapUID:                   "yQ4ktCXu3SAxyRx9gar8hj7kVBb",
			Name:                     "Spring 2024 - 01",
			SilverScore:              35000,
			Submitter:                "d2372a08-a8a1-46cb-97fb-23a161d85ad0",
			FileURL:                  "https://core.trackmania.nadeo.live/storageObjects/69ff7efb-010e-43db-8233-11d119aa0499",
			ThumbnailURL:             "https://core.trackmania.nadeo.live/storageObjects/ea9128eb-dabb-4b08-b9d1-37adb25b41b2.jpg",
		},
	}
	m.DB.Save(tracks)
	mappackTracks := []models.MappackTrack{
		{
			MappackID: "mappack-beginner",
			TrackID:   tracks[0].ID,
			Tier:      "A",
		},
		{
			MappackID: "mappack-advanced",
			TrackID:   tracks[0].ID,
			Tier:      "B",
		},
		{
			MappackID: "mappack-pro",
			TrackID:   tracks[0].ID,
			Tier:      "C",
		},
	}
	return m.DB.Save(mappackTracks).Error
}
