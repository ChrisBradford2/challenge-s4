package models

type Hackathon struct {
	Base
	Name            string `json:"name" gorm:"not null"`
	Description     string `json:"description"`
	Location        string `json:"location"`
	MaxParticipants int    `json:"maxParticipants"`
	CreatedBy       User   `gorm:"foreignKey:ID"`
	StartDate       string
	EndDate         string
}

type HackathonCreate struct {
	Name            string `json:"name" example:"Hackathon de Paris"`
	Description     string `json:"description" example:"Un événement pour les développeurs"`
	Location        string `json:"location" example:"Paris"`
	MaxParticipants int    `json:"maxParticipants" example:"100"`
}
