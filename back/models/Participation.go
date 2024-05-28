package models

import "gorm.io/gorm"

type Participation struct {
	gorm.Model
	UserID      uint `gorm:"primaryKey"`    // Foreign key referencing User.ID
	HackathonID uint `gorm:"primaryKey"`    // Foreign key referencing Hackathon.ID
	IsOrganizer bool `gorm:"default:false"` // Indicates if the user is an organizer of the hackathon
}
