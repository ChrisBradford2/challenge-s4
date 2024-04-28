package models

type Evaluation struct {
	Base
	TeamID  uint // Foreign key referencing Team.ID
	Team    Team // Belongs to Team
	Note    uint8
	Comment string
	UserID  uint `gorm:"column:author_id"` // Foreign key referencing User.ID, column name as "author_id"
	Author  User `gorm:"foreignKey:UserID"`
}
