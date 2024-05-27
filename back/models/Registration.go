package models

type Registration struct {
	Base
	UserID  uint // Foreign key referencing User.ID
	User    User // Belongs to User
	TeamID  uint // Foreign key referencing Team.ID
	Team    Team // Belongs to Team
	IsValid bool `json:"is_valid" gorm:"default:false" example:"false"`
}

type RegistrationCreate struct {
	UserID  uint `json:"user_id" binding:"required" example:"1"`
	TeamID  uint `json:"team_id" binding:"required" example:"1"`
	IsValid bool `json:"is_valid" example:"false"`
}

type RegistrationUpdate struct {
	IsValid bool `json:"is_valid" example:"true"`
}
