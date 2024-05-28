package seeders

import (
	"challenges4/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

// SeedHackathons seeds the database with hackathon data.
func SeedHackathons(db *gorm.DB) error {
	hackathons := []models.Hackathon{
		{
			Name:            "Hackathon de Paris",
			Description:     "Un événement pour les développeurs",
			Address:         "Paris",
			MaxParticipants: 100,
			StartDate:       "2021-01-01",
			EndDate:         "2021-01-02",
		},
		{
			Name:            "Hackathon de Londres",
			Description:     "Un événement pour les développeurs à Londres",
			Address:         "Londres",
			MaxParticipants: 150,
			StartDate:       "2021-03-15",
			EndDate:         "2021-03-17",
		},
		{
			Name:            "Hackathon de New York",
			Description:     "Un événement pour les développeurs à New York",
			Address:         "New York",
			MaxParticipants: 200,
			StartDate:       "2021-06-10",
			EndDate:         "2021-06-12",
		},
	}

	for _, hackathon := range hackathons {
		var existingHackathon models.Hackathon
		result := db.Where("name = ?", hackathon.Name).First(&existingHackathon)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		if result.RowsAffected == 0 {
			if err := db.Create(&hackathon).Error; err != nil {
				return err
			}
			log.Printf("Created hackathon with name %s", hackathon.Name)
		} else {
			log.Printf("Hackathon with name %s already exists", hackathon.Name)
		}
	}
	return nil
}
