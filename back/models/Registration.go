package models

type Registration struct {
	Base
	UserID uint // Foreign key referencing User.ID
	User   User // Belongs to User
	TeamID uint // Foreign key referencing Team.ID
	Team   Team // Belongs to Team
	Status string
}
