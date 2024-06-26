package controllers

import (
	"challenges4/config"
	"challenges4/models"
	"challenges4/services"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 500 {object} string "Internal server error"
// @Router /hackathons [post]
func CreateHackathon(c *gin.Context, db *gorm.DB) {
	var hackathonCreate models.HackathonCreate

	if err := c.ShouldBindJSON(&hackathonCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	hackathon := models.Hackathon{
		Name:                   hackathonCreate.Name,
		Description:            hackathonCreate.Description,
		Address:                hackathonCreate.Address,
		Location:               hackathonCreate.Location,
		MaxParticipants:        hackathonCreate.MaxParticipants,
		MaxParticipantsPerTeam: hackathonCreate.MaxParticipantsPerTeam,
		StartDate:              hackathonCreate.StartDate,
		EndDate:                hackathonCreate.EndDate,
		CreatedByID:            &userID,
	}

	if result := db.Create(&hackathon); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
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

	if !config.HasRole(&user, config.RoleOrganizer) {
		user.Roles |= config.RoleOrganizer
		if updateRoleResult := db.Save(&user); updateRoleResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": updateRoleResult.Error.Error()})
			return
		}
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

// GetHackathonsByUser godoc
// @Summary Get hackathons created by the user
// @Description Get hackathons created by the user
// @Tags Hackathons
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.Hackathon "Successfully retrieved list of hackathons"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /hackathons/user [get]
func GetHackathonsByUser(c *gin.Context, db *gorm.DB) {
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

	var hackathons []models.Hackathon
	if result := db.Where("created_by_id = ?", userID).Find(&hackathons); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hackathons})
}

// GetHackathon godoc
// @Summary Get a single Hackathon
// @Description Get a single Hackathon
// @Tags Hackathons
// @Produce  json
// @Param id path int true "Hackathon ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Hackathon "Successfully retrieved Hackathon"
// @Failure 404 {object} string "Hackathon not found"
// @Router /hackathons/{id} [get]
func GetHackathon(c *gin.Context, db *gorm.DB) {
	hackathonID := c.Param("id")
	var hackathon models.Hackathon

	if err := db.Preload("Teams.Users").First(&hackathon, hackathonID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hackathon})
}

// UpdateHackathon godoc
// @Summary Update a Hackathon
// @Description Update a Hackathon
// @Tags Hackathons
// @Accept  json
// @Produce  json
// @Param id path int true "Hackathon ID"
// @Param hackathon body models.HackathonCreate true "Hackathon object"
// @Security ApiKeyAuth
// @Success 200 {object} models.Hackathon "Successfully updated Hackathon"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Hackathon not found"
// @Failure 500 {object} string "Internal server error"
// @Router /hackathons/{id} [put]
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

// GetTeamsByHackathon godoc
// @Summary Get all teams for a specific hackathon
// @Description Get all teams for a specific hackathon
// @Tags Teams
// @Produce json
// @Security ApiKeyAuth
// @Param hackathonId path int true "Hackathon ID"
// @Success 200 {array} models.Team
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /hackathons/{hackathonId}/teams [get]
func GetTeamsByHackathon(c *gin.Context, db *gorm.DB) {
	hackathonIdStr := c.Param("id")

	// Convertir l'ID du hackathon en entier
	hackathonId, err := strconv.Atoi(hackathonIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

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

	// Vérifiez si l'utilisateur est l'organisateur du hackathon
	var hackathon models.Hackathon
	if err := db.First(&hackathon, hackathonId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

	if hackathon.CreatedByID != nil && *hackathon.CreatedByID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: You are not the organizer of this hackathon"})
		return
	}

	var teams []models.Team
	if err := db.Preload("Users").Preload("Hackathon").Preload("Submission").Where("hackathon_id = ?", hackathonId).Find(&teams).Error; err != nil { // Preload users, hackathon, and submission
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teams"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teams})
}
