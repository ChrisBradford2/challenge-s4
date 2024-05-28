package models

import (
	"fmt"
	"google.golang.org/genproto/googleapis/type/date"
	"gorm.io/gorm"
)

type Hackathon struct {
	Base
	Name            string          `json:"name" gorm:"not null"`
	Description     string          `json:"description"`
	Address         string          `json:"address"`
	Longitude       float64         `json:"longitude"`
	Latitude        float64         `json:"latitude"`
	MaxParticipants int             `json:"max_participants"`
	CreatedByID     *uint           `json:"created_by_id"`
	CreatedBy       *User           `gorm:"foreignKey:CreatedByID"`
	Teams           []Team          `json:"teams" gorm:"foreignKey:HackathonID"`
	StartDate       string          `json:"start_date"`
	EndDate         string          `json:"end_date"`
	IsActive        bool            `json:"is_active" gorm:"default:0"`
	HackathonFiles  []File          `json:"hackathon_files" gorm:"foreignKey:HackathonID"`
	Participations  []Participation `gorm:"foreignKey:HackathonID"` // Many-to-many relationship with User through Participation
	NbOfTeams       int             `json:"nb_of_teams" gorm:"default:0"`
}

type HackathonCreate struct {
	Name            string    `json:"name" example:"Hackathon de Paris"`
	Description     string    `json:"description" example:"Un événement pour les développeurs"`
	Location        string    `json:"location" example:"Paris"`
	MaxParticipants int       `json:"max_participants" example:"100"`
	StartDate       date.Date `json:"start_date" example:"2021-01-01"`
	EndDate         date.Date `json:"end_date" example:"2021-01-02"`
	NbOfTeams       int       `json:"nb_of_teams" example:"0"`
	Address         string    `json:"address" example:"Paris"`
	Longitude       float64   `json:"longitude" example:"0.0"`
	Latitude        float64   `json:"latitude" example:"0.0"`
}

type ParticipationFilter struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	SkillID  uint   `json:"skill_id"`
}

func (h *Hackathon) AfterCreate(tx *gorm.DB) error {
	for i := 1; i <= h.NbOfTeams; i++ {
		team := &Team{
			Name:        fmt.Sprintf("Team %d", i),
			HackathonID: &h.ID,
		}
		if err := tx.Create(team).Error; err != nil {
			return err
		}
	}
	return nil
}
