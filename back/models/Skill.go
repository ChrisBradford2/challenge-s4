package models

type Skill struct {
	Base
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Users       []User `gorm:"many2many:user_skills;"` // Many-to-many relationship with User
}
