package models

type Notification struct {
	Base
	HackathonID uint      `json:"hackathon_id"`           // Foreign key referencing Hackathon.ID
	Hackathon   Hackathon `gorm:"foreignKey:HackathonID"` // Belongs to Hackathon
	CreatedByID *uint     `json:"created_by_id"`
	CreatedBy   *User     `gorm:"foreignKey:CreatedByID"`
	Message     string    `json:"message"`
}
