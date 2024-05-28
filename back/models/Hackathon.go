package models

type Hackathon struct {
	Base
	Name            string          `json:"name" gorm:"not null" example:"Hackathon de Paris"`
	Description     string          `json:"description" gorm:"not null" example:"Un événement pour les développeurs"`
	Location        string          `json:"location" gorm:"not null" example:"Paris"`
	MaxParticipants int             `json:"max_participants" gorm:"not null" example:"100"`
	CreatedByID     *uint           `json:"created_by_id"`
	CreatedBy       *User           `gorm:"foreignKey:CreatedByID" json:"created_by"`
	Teams           []Team          `gorm:"many2many:hackathon_teams;"`
	StartDate       string          `json:"start_date" gorm:"not null" example:"2021-01-01"`
	EndDate         string          `json:"end_date" gorm:"not null" example:"2021-01-02"`
	IsActive        bool            `json:"is_active" gorm:"default:0"`
	HackathonFiles  []File          `json:"hackathon_files" gorm:"foreignKey:HackathonID"`
	Participations  []Participation `gorm:"foreignKey:HackathonID"` // Many-to-many relationship with User through Participation
}

type HackathonCreate struct {
	Name            string `json:"name" example:"Hackathon de Paris"`
	Description     string `json:"description" example:"Un événement pour les développeurs"`
	Location        string `json:"location" example:"Paris"`
	MaxParticipants int    `json:"max_participants" example:"100"`
	StartDate       string `json:"start_date" example:"2021-01-01"`
	EndDate         string `json:"end_date" example:"2021-01-02"`
}
