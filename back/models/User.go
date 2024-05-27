package models

import "gorm.io/gorm"

type User struct {
	Base
	Username       string          `gorm:"unique" json:"username" binding:"required" example:"jdoe"`
	LastName       string          `json:"last_name" binding:"required" example:"Doe"`
	FirstName      string          `json:"first_name" binding:"required" example:"John"`
	Email          string          `gorm:"unique" json:"email" binding:"required" example:"john.doe@exmple.com"`
	Password       string          `gorm:"not null" json:"password" binding:"required" example:"password"`
	ProfilePicture string          `json:"profile_picture"`
	TeamID         *uint           `json:"team_id" gorm:"column:team_id"`           // Foreign key referencing Team.ID
	Team           *Team           `json:"team,omitempty" gorm:"foreignKey:TeamID"` // Belongs to Team
	Roles          uint8           `json:"roles" example:"0"`                       // 0 = user, 2 = organizer, 4 = admin
	CreatedByID    *uint           `json:"created_by_id"`
	CreatedBy      *User           `gorm:"foreignKey:CreatedByID"`
	Participations []Participation `gorm:"foreignKey:UserID"` // Many-to-many relationship with Hackathon through Participation
}

type UserRegister struct {
	Username       string `json:"username" binding:"required" example:"jdoe"`
	LastName       string `json:"last_name" binding:"required" example:"Doe"`
	FirstName      string `json:"first_name" binding:"required" example:"John"`
	ProfilePicture string `json:"profile_picture"`
	Email          string `json:"email" binding:"required" example:"john.doe@exmple.com"`
	Password       string `json:"password" binding:"required" example:"password"`
}

type UserRegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required" example:"john.doe@exmple.com"`
	Password string `json:"password" binding:"required" example:"password"`
}

// PublicUser omits sensitive data from user model
type PublicUser struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
