package models

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Hackathon struct {
	Base
	Name            string `json:"name" gorm:"not null"`
	Description     string `json:"description"`
	Location        string `json:"location"`
	MaxParticipants int    `json:"maxParticipants"`
	CreatedBy       User   `gorm:"foreignKey:ID"`
	Teams           []Team `gorm:"many2many:hackathon_teams;"`
	StartDate       string
	EndDate         string
}

type HackathonCreate struct {
	Name            string `json:"name" example:"Hackathon de Paris"`
	Description     string `json:"description" example:"Un événement pour les développeurs"`
	Location        string `json:"location" example:"Paris"`
	MaxParticipants int    `json:"maxParticipants" example:"100"`
}
