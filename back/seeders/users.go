package seeders

import (
	"challenges4/config"
	"challenges4/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

// SeedUsers seeds the database with user data.
func SeedUsers(db *gorm.DB) error {
	users := []models.User{
		{Username: "user1", LastName: "Doe", FirstName: "John", Email: "johndoe@gmail.com", Password: "$2a$14$M0HeLvAkI0xqB0CUUR4tEOAXnsEH8lv4h55aaU2vDZkzuA4zrGmsO", Roles: config.RoleUser},
		{Username: "user2", LastName: "Smith", FirstName: "Jane", Email: "janesmith@gmail.com", Password: "$2a$14$vVHG/vtq8qYUfxe2CNhDlelk6DGWQHYaPI2v4rW/an5gWiOZQi8LK", Roles: config.RoleUser},
		{Username: "user3", LastName: "Johnson", FirstName: "Alice", Email: "testOrg@gmail.com", Password: "$2a$14$Qi.XxbNIKjVbLsgEmxAMX.kX5Y.zEUfJVGG1N.HYSB2W1Ol54OzLa", Roles: config.RoleOrganizer},
		{Username: "admin", LastName: "Admin", FirstName: "admin", Email: "test@gmail.com", Password: "$2a$14$Qi.XxbNIKjVbLsgEmxAMX.kX5Y.zEUfJVGG1N.HYSB2W1Ol54OzLa", Roles: config.RoleAdmin},
	}

	for _, user := range users {
		var existingUser models.User
		result := db.Where("email = ?", user.Email).First(&existingUser)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		if result.RowsAffected == 0 {
			if err := db.Create(&user).Error; err != nil {
				return err
			}
			log.Printf("Created user with email %s", user.Email)
		} else {
			log.Printf("User with email %s already exists", user.Email)
		}
	}
	return nil
}
