package models

type Team struct {
	Base
	Name        string     `gorm:"unique" json:"name" binding:"required" example:"Team 1"`
	Users       []User     `json:"users"` // Has many Users
	HackathonID *uint      `json:"hackathon_id"`
	Hackathon   *Hackathon `gorm:"foreignKey:HackathonID"`
	NbOfMembers int        `json:"nbOfMembers"`
}
