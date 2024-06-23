package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Hackathon struct {
	Base
	Name                   string          `json:"name" gorm:"not null"`
	Description            string          `json:"description"`
	Address                string          `json:"address"`
	Location               string          `json:"location" gorm:"not null"`
	MaxParticipants        int             `json:"max_participants"`
	MaxParticipantsPerTeam int             `json:"max_participants_per_team"`
	CreatedByID            *uint           `json:"created_by_id"`
	CreatedBy              *User           `gorm:"foreignKey:CreatedByID"`
	Teams                  []Team          `json:"teams" gorm:"foreignKey:HackathonID"`
	StartDate              string          `json:"start_date"`
	EndDate                string          `json:"end_date"`
	IsActive               bool            `json:"is_active" gorm:"default:0"`
	HackathonFiles         []File          `json:"hackathon_files" gorm:"foreignKey:HackathonID"`
	Participations         []Participation `gorm:"foreignKey:HackathonID"` // Many-to-many relationship with User through Participation
	NbOfTeams              int             `json:"nb_of_teams" gorm:"default:0"`
}

type HackathonCreate struct {
	Name                   string `json:"name" example:"Hackathon de Paris"`
	Description            string `json:"description" example:"Un événement pour les développeurs"`
	Location               string `json:"location" example:"Paris"`
	MaxParticipants        int    `json:"max_participants" example:"100"`
	MaxParticipantsPerTeam int    `json:"max_participants_per_team"`
	StartDate              string `json:"start_date" example:"2021-01-01"`
	EndDate                string `json:"end_date" example:"2021-01-02"`
	Address                string `json:"address" example:"Paris"`
}

type ParticipationFilter struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	SkillID  uint   `json:"skill_id"`
}

func (h *Hackathon) AfterCreate(db *gorm.DB) (err error) {
	if h.MaxParticipants <= 0 || h.MaxParticipantsPerTeam <= 0 {
		return nil
	}

	// Calculate the number of teams needed
	nbOfTeams := (h.MaxParticipants + h.MaxParticipantsPerTeam - 1) / h.MaxParticipantsPerTeam
	h.NbOfTeams = nbOfTeams

	for i := 0; i < nbOfTeams; i++ {
		team := Team{
			Name:        fmt.Sprintf("Team %d (Hackathon %d)", i+1, h.ID),
			HackathonID: &h.ID,
			NbOfMembers: h.MaxParticipantsPerTeam, // Set the number of members per team
		}
		if result := db.Create(&team); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
