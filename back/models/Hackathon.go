package models

type Hackathon struct {
	Base
	Name            string `json:"name" gorm:"not null"`
	Description     string `json:"description"`
	Location        string `json:"location"`
	MaxParticipants int    `json:"max_participants"`
	CreatedByID     *uint  `json:"created_by_id"`
	CreatedBy       *User  `gorm:"foreignKey:CreatedByID"`
	Teams           []Team `gorm:"many2many:hackathon_teams;"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	HackathonFiles  []File `json:"hackathon_files" gorm:"foreignKey:HackathonID"`
}

type HackathonCreate struct {
	Name            string `json:"name" example:"Hackathon de Paris"`
	Description     string `json:"description" example:"Un événement pour les développeurs"`
	Location        string `json:"location" example:"Paris"`
	MaxParticipants int    `json:"max_participants" example:"100"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
}
