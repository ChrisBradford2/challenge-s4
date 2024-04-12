package models

import "gorm.io/gorm"

type Hackathon struct {
	gorm.Model
	ID        uint `gorm:"not null"`
	Nom       string
	Location  string
	CreatedBy User `gorm:"foreignKey:ID"`
	Teams     []Team
	StartDate string
	EndDate   string
}
