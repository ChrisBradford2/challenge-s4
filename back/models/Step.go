package models

import (
	"time"
)

type Step struct {
	Base
	Position     uint       `json:"position" example:"1"`
	Title        string     `json:"title" example:"Step 1"`
	HackathonID  *uint      `json:"hackathon_id" example:"1"`
	Hackathon    *Hackathon `gorm:"foreignKey:HackathonID"`
	Status       string     `json:"status" example:"done"`
	DeadLineDate time.Time  `json:"dead_line_date" example:"2025-12-31T23:59:59Z"`
}

type StepCreate struct {
	Position     uint      `json:"position" example:"1"`
	Title        string    `json:"title" example:"Step 1"`
	HackathonID  *uint     `json:"hackathon_id" example:"1"`
	DeadLineDate time.Time `json:"dead_line_date" example:"2025-12-31T23:59:59Z"`
}

type StepUpdate struct {
	Title        *string   `json:"title" example:"Step 1"`
	Status       *string   `json:"status" example:"done"`
	DeadLineDate time.Time `json:"dead_line_date" example:"2025-12-31T23:59:59Z"`
}

type StepPublic struct {
	ID           uint      `json:"id" example:"1"`
	Position     uint      `json:"position" example:"1"`
	Title        string    `json:"title" example:"Step 1"`
	HackathonID  *uint     `json:"hackathon_id" example:"1"`
	Status       string    `json:"status" example:"done"`
	DeadLineDate time.Time `json:"dead_line_date" example:"2025-12-31T23:59:59Z"`
}
