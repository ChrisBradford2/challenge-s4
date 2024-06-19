package middleware

import (
	"challenges4/models"
	"fmt"
	"gorm.io/gorm"
)

func CreateTeamsForHackathon(db *gorm.DB, hackathon *models.Hackathon) error {
	if hackathon.MaxParticipants <= 0 || hackathon.MaxParticipantsPerTeam <= 0 {
		return nil
	}

	// Calculate the number of teams needed
	nbOfTeams := (hackathon.MaxParticipants + hackathon.MaxParticipantsPerTeam - 1) / hackathon.MaxParticipantsPerTeam
	hackathon.NbOfTeams = nbOfTeams

	for i := 0; i < nbOfTeams; i++ {
		team := models.Team{
			Name:        fmt.Sprintf("Team %d (Hackathon %d)", i+1, hackathon.ID),
			HackathonID: &hackathon.ID,
		}
		if result := db.Create(&team); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
