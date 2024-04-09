package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name          string         `gorm:"unique"`
	Users         []User         // Has many Users
	Registrations []Registration // Has many Registrations
}
