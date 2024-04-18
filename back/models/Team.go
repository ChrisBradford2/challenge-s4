package models

type Team struct {
	Base
	Name          string         `gorm:"unique"`
	Users         []User         // Has many Users
	Registrations []Registration // Has many Registrations
}
