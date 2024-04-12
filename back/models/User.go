package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	LastName  string
	FirstName string
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
	TeamID    uint   // Foreign key referencing Team.ID
	Team      Team   // Belongs to Team
	Roles     uint8  // 0 = user, 2 = organizer, 4 = admin
}
