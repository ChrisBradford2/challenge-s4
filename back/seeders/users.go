package seeders

import (
	"challenges4/config"
	"challenges4/models"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	users := []models.User{
		{Username: "user1", LastName: "Doe", FirstName: "John", Email: "johndoe@gmail.com", Password: "$2a$14$M0HeLvAkI0xqB0CUUR4tEOAXnsEH8lv4h55aaU2vDZkzuA4zrGmsO", Roles: config.RoleUser},
		{Username: "user2", LastName: "Smith", FirstName: "Jane", Email: "janesmith@gmail.com", Password: "$2a$14$vVHG/vtq8qYUfxe2CNhDlelk6DGWQHYaPI2v4rW/an5gWiOZQi8LK", Roles: config.RoleUser},
		{Username: "admin", LastName: "Admin", FirstName: "admin", Email: "test@gmail.com", Password: "$2a$14$Qi.XxbNIKjVbLsgEmxAMX.kX5Y.zEUfJVGG1N.HYSB2W1Ol54OzLa", Roles: config.RoleAdmin},
	}

	return db.Create(&users).Error
}
