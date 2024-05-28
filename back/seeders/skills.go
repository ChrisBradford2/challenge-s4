package seeders

import (
	"challenges4/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

func SeedSkills(db *gorm.DB) error {
	skills := []models.Skill{
		{Name: "FrontEnd dev", Description: "Frontend development skills"},
		{Name: "BackEnd dev", Description: "Backend development skills"},
		{Name: "Marketing", Description: "Marketing skills"},
		{Name: "UI/UX Designer", Description: "UI/UX design skills"},
	}

	for _, skill := range skills {
		var existingSkill models.Skill
		result := db.Where("name = ?", skill.Name).First(&existingSkill)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		if result.RowsAffected == 0 {
			if err := db.Create(&skill).Error; err != nil {
				return err
			}
			log.Printf("Created skill with name %s", skill.Name)
		} else {
			log.Printf("Skill with name %s already exists", skill.Name)
		}
	}
	return nil
}
