package controllers

import (
	"challenges4/models"
	"challenges4/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// CreateHackathon godoc
// @Summary Create a new Hackathon
// @Description Create a new Hackathon
// @Tags Hackathons
// @Accept  json
// @Produce  json
// @Param hackathon body models.HackathonCreate true "Hackathon object"
// @Security ApiKeyAuth
// @Success 201 {object} models.Hackathon "Successfully created Hackathon"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /hackathons [post]
func CreateHackathon(c *gin.Context, db *gorm.DB) {
	var hackathon models.Hackathon

	if err := c.ShouldBindJSON(&hackathon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&hackathon); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	participation := models.Participation{
		HackathonID: hackathon.ID,
		UserID:      user.ID,
		IsOrganizer: true,
	}

	if createParticipationResult := db.Create(&participation); createParticipationResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createParticipationResult.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": hackathon})
}

// GetHackathons godoc
// @Summary Get all Hackathons
// @Description Get all Hackathons
// @Tags Hackathons
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.Hackathon
// @Router /hackathons [get]
func GetHackathons(c *gin.Context, db *gorm.DB) {
	var hackathons []models.Hackathon

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	isActive := c.Query("active")
	query := db.Model(&models.Hackathon{})

	/*
		if user.Roles != config.RoleAdmin {
			query = query.Joins("JOIN participations ON hackathons.id = participations.hackathon_id").
				Where("participations.user_id = ?", user.ID)
		}
	*/

	if isActive != "" {
		query = db.Where("is_active = ?", isActive)
	}

	if result := query.Find(&hackathons); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hackathons})
}

func SearchHackathons(c *gin.Context, db *gorm.DB) {
	var searchCriteria struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
		Distance  float64 `json:"distance"`
	}

	if err := c.ShouldBindJSON(&searchCriteria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := db.Model(&models.Hackathon{})

	if searchCriteria.Longitude != 0 && searchCriteria.Latitude != 0 {
		query = query.Select("*, ST_DistanceSphere(location, ST_MakePoint(?, ?)) as distance", searchCriteria.Longitude, searchCriteria.Latitude).
			Order("distance")

		if searchCriteria.Distance != 0 {
			query = query.Where("ST_DistanceSphere(location, ST_MakePoint(?, ?)) <= ?", searchCriteria.Longitude, searchCriteria.Latitude, searchCriteria.Distance)
		}
	}

	var hackathons []models.Hackathon
	if result := query.Find(&hackathons); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hackathons})
}

func GetHackathon(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var hackathon models.Hackathon
	if result := db.First(&hackathon, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hackathon})
}

func UpdateHackathon(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var hackathon models.Hackathon
	if err := db.First(&hackathon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

	var input models.Hackathon
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var currentHackathon models.Hackathon
	if err := db.First(&currentHackathon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

	if currentHackathon.NbOfTeams != input.NbOfTeams {
		if currentHackathon.IsActive {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot change number of teams for an active hackathon"})
			return
		} else {
			// Unlink and delete the excess teams
			if input.NbOfTeams < currentHackathon.NbOfTeams {
				var teams []models.Team
				if err := db.Where("hackathon_id = ?", currentHackathon.ID).Limit(currentHackathon.NbOfTeams - input.NbOfTeams).Find(&teams).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teams"})
					return
				}

				for _, team := range teams {
					if err := db.Model(&team).Update("hackathon_id", nil).Error; err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlink team from hackathon"})
						return
					}
					if err := db.Delete(&team).Error; err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete team"})
						return
					}
				}
			} else {
				for i := currentHackathon.NbOfTeams + 1; i <= input.NbOfTeams; i++ {
					team := &models.Team{
						Name:        fmt.Sprintf("Team %d of hackathon: %s", i, currentHackathon.Name),
						HackathonID: &hackathon.ID,
					}
					if err := db.Create(team).Error; err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
						return
					}
				}
			}
		}
	}

	db.Model(&hackathon).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": hackathon})
}

// DeleteHackathon godoc
// @Summary Delete a Hackathon
// @Description Delete a Hackathon
// @Tags Hackathons
// @Produce  json
// @Param id path int true "Hackathon ID"
// @Security ApiKeyAuth
// @Success 200 {object} bool "Successfully deleted Hackathon"
// @Router /hackathons/{id} [delete]
func DeleteHackathon(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&models.Hackathon{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// HackathonRegister godoc
// @Summary Register for a Hackathon
// @Description Register for a Hackathon
// @Tags Hackathons
// @Produce  json
// @Param id path int true "Hackathon ID"
// @Security ApiKeyAuth
// @Success 200 {object} bool "Successfully registered for Hackathon"
// @Router /hackathons/{id}/register [post]
func HackathonRegister(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var hackathon models.Hackathon
	if err := db.First(&hackathon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var participation models.Participation
	participation.HackathonID = hackathon.ID
	participation.UserID = user.ID

	if result := db.Create(&participation); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Successfully registered for hackathon"})

}

func isHackathonOrganizer(db *gorm.DB, hackathonID uint, userID uint) bool {
	var participation models.Participation
	if result := db.Where("hackathon_id = ? AND user_id = ?", hackathonID, userID).First(&participation); result.Error != nil {
		return false
	}

	return participation.IsOrganizer
}

func SearchParticipants(c *gin.Context, db *gorm.DB) ([]models.User, error) {

	hackathonID := c.Param("id")

	var participantFilter models.ParticipationFilter
	if err := c.ShouldBindJSON(&participantFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	var users []models.User
	queryBuilder := db.Joins("JOIN participations ON users.id = participations.user_id")

	if participantFilter.SkillID != 0 {
		queryBuilder = queryBuilder.Joins("JOIN user_skills ON users.id = user_skills.user_id")
	}

	queryBuilder = queryBuilder.Where("participations.hackathon_id = ?", hackathonID)

	if participantFilter.Username != "" {
		queryBuilder = queryBuilder.Where("users.username LIKE ?", "%"+participantFilter.Username+"%")
	}

	if participantFilter.Email != "" {
		queryBuilder = queryBuilder.Where("users.email LIKE ?", "%"+participantFilter.Email+"%")
	}

	if participantFilter.SkillID != 0 {
		queryBuilder = queryBuilder.Where("user_skills.skill_id = ?", participantFilter.SkillID)
	}

	if result := queryBuilder.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
