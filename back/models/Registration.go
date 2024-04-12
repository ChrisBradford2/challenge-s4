package models

import "gorm.io/gorm"

type Registration struct {
	gorm.Model
	UserID uint // Foreign key referencing User.ID
	User   User // Belongs to User
	TeamID uint // Foreign key referencing Team.ID
	Team   Team // Belongs to Team
	Status string
}
