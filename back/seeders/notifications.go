package seeders

import (
	"challenges4/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

// SeedNotifications seeds the database with notification data.
func SeedNotifications(db *gorm.DB) error {
	// Example IDs; replace these with valid IDs from your database
	hackathonIDs := []uint{1, 2, 3} // Replace with actual Hackathon IDs
	userID := uint(1)               // Assume a user with ID 1 exists

	notifications := []models.Notification{
		{
			Message:     "Hackathon de Paris",
			HackathonID: hackathonIDs[0],
			CreatedByID: &userID,
		},
		{
			Message:     "Hackathon de Londres",
			HackathonID: hackathonIDs[1],
			CreatedByID: &userID,
		},
		{
			Message:     "Hackathon de New York",
			HackathonID: hackathonIDs[2],
			CreatedByID: &userID,
		},
	}

	for _, notification := range notifications {
		var existingNotification models.Notification
		result := db.Where("message = ?", notification.Message).First(&existingNotification)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		if result.RowsAffected == 0 {
			if err := db.Create(&notification).Error; err != nil {
				return err
			}
			log.Printf("Created notification with message %s", notification.Message)
		} else {
			log.Printf("Notification with message %s already exists", notification.Message)
		}
	}
	return nil
}
