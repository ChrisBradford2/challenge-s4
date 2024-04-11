package models

import (
	"time"
)

type Base struct {
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
}

type HackathonCreate struct {
	Name            string `json:"name" example:"Hackathon de Paris"`
	Description     string `json:"description" example:"Un événement pour les développeurs"`
	Location        string `json:"location" example:"Paris"`
	MaxParticipants int    `json:"maxParticipants" example:"100"`
}
