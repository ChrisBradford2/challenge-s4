package models

type User struct {
	Base
	Username  string `gorm:"unique" json:"username" binding:"required" example:"jdoe"`
	LastName  string `json:"last_name" binding:"required" example:"Doe"`
	FirstName string `json:"first_name" binding:"required" example:"John"`
	Email     string `gorm:"unique" json:"email" binding:"required" example:"john.doe@exmple.com"`
	Password  string `gorm:"not null" json:"password" binding:"required" example:"password"`
	TeamID    uint   `json:"team_id"`           // Foreign key referencing Team.ID
	Team      Team   `json:"team"`              // Belongs to Team
	Roles     uint8  `json:"roles" example:"0"` // 0 = user, 2 = organizer, 4 = admin
}
