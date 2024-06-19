package middleware

import (
	"challenges4/models"
	"fmt"
	"gorm.io/gorm"
)

func CreateTeamsForHackathon(db *gorm.DB, hackathon *models.Hackathon) error {
	if hackathon.NbOfTeams <= 0 || hackathon.MaxParticipants <= 0 {
		return nil
	}

	teamSize := hackathon.MaxParticipants / hackathon.NbOfTeams
	for i := 0; i < hackathon.NbOfTeams; i++ {
		teamName := fmt.Sprintf("Team %d (Hackathon %d)", i+1, hackathon.ID)

		// Check if the team name already exists
		var existingTeam models.Team
		for db.Where("name = ?", teamName).First(&existingTeam).RowsAffected > 0 {
			teamName = fmt.Sprintf("Team %d (Hackathon %d) - %d", i+1, hackathon.ID, i)
		}

		team := models.Team{
			Name:        teamName,
			HackathonID: &hackathon.ID,
			NbOfMembers: teamSize,
		}
		if result := db.Create(&team); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
