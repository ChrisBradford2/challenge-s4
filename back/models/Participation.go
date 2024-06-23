package models

type Participation struct {
	Base
	UserID      uint  `gorm:"primaryKey"`    // Foreign key referencing User.ID
	HackathonID uint  `gorm:"primaryKey"`    // Foreign key referencing Hackathon.ID
	IsOrganizer bool  `gorm:"default:false"` // Indicates if the user is an organizer of the hackathon
	TeamID      *uint `json:"team_id"`
}
