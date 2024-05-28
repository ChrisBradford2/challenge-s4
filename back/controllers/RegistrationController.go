package controllers

import (
	"challenges4/config"
	"challenges4/models"
	"challenges4/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RegistrationController struct {
	DB *gorm.DB
}

func NewRegistrationController(db *gorm.DB) *RegistrationController {
	return &RegistrationController{DB: db}
}

// CreateRegistration godoc
// @Summary Create a new registration
// @Description Create a new registration
// @Tags registrations
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Registration "Successfully retrieved registration"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /registrations [post]
func (ctrl *RegistrationController) CreateRegistration(c *gin.Context) {
	var registrationCreate models.RegistrationCreate
	if err := c.ShouldBindJSON(&registrationCreate); err != nil {
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
	if err := ctrl.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	registration := models.RegistrationCreate{
		UserID:  userID,
		TeamID:  registrationCreate.TeamID,
		IsValid: user.Roles == config.RoleOrganizer || user.Roles == config.RoleAdmin,
	}

	if result := ctrl.DB.Create(&registration); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": registration})
}

// UpdateRegistration godoc
// @Summary Update a registration
// @Description Update a registration by ID
// @Tags registrations
// @Accept  json
// @Produce  json
// @Param id path int true "Registration ID"
// @Param registration body models.RegistrationUpdate true "Registration update object"
// @Security ApiKeyAuth
// @Success 200 {object} models.Registration "Successfully updated registration"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Registration not found"
// @Failure 500 {object} string "Internal server error"
// @Router /registrations/{id} [patch]
func (ctrl *RegistrationController) UpdateRegistration(c *gin.Context) {
	registrationID := c.Param("id")
	var existingRegistration models.Registration
	if err := ctrl.DB.First(&existingRegistration, registrationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Registration not found"})
		return
	}

	var registrationUpdate models.RegistrationUpdate
	if err := c.ShouldBindJSON(&registrationUpdate); err != nil {
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
	if err := ctrl.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var team models.Team
	if err := ctrl.DB.First(&team, existingRegistration.TeamID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	if team.HackathonID == nil || !isHackathonOrganizer(ctrl.DB, *team.HackathonID, userID) && user.Roles != config.RoleAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Unauthorized to update registration"})
		return
	}

	if err := ctrl.DB.Model(&existingRegistration).Updates(registrationUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update registration"})
		return
	}

	if registrationUpdate.IsValid {
		// Add the user to the team's users association
		if err := ctrl.DB.Model(&team).Association("Users").Append(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to the team"})
			return
		}
	} else {
		// Remove the user from the team's users association
		if err := ctrl.DB.Model(&team).Association("Users").Delete(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove user from the team"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": existingRegistration})
}

// GetRegistrations godoc
// @Summary Get registrations with optional filters
// @Description Retrieves registrations from the database with optional filtering based on teamid, user, and status
// @Tags registrations
// @Produce  json
// @Security ApiKeyAuth
// @Param teamid query int false "Filter by team ID"
// @Param user query string false "Filter by user"
// @Param status query string false "Filter by status"
// @Success 200 {array} models.Registration
// @Router /registrations [get]
func (ctrl *RegistrationController) GetRegistrations(c *gin.Context) {
	var registrations []models.Registration

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
	if err := ctrl.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Roles == config.RoleAdmin {
		query := ctrl.DB
		applyFilters(c, query)

		if result := query.Find(&registrations); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	} else {
		var hackathons []models.Hackathon
		if result := ctrl.DB.Where("organizer_id = ?", userID).Find(&hackathons); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		var teams []models.Team
		if result := ctrl.DB.Find(&teams).Where("hackathon_id IN ?", hackathons); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		query := ctrl.DB.Where("team_id IN ?", teams)
		applyFilters(c, query)

		if result := query.Find(&registrations); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": registrations})
}

func applyFilters(c *gin.Context, query *gorm.DB) {
	// Get query parameters
	teamID := c.Query("teamId")
	targetUser := c.Query("user")
	status := c.Query("status")

	// Apply filters based on query parameters
	if teamID != "" {
		tid, err := strconv.Atoi(teamID)
		if err == nil {
			query = query.Where("team_id = ?", tid)
		}
	}
	if targetUser != "" {
		query = query.Where("user = ?", targetUser)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
}

// GetRegistration godoc
// @Summary Get a registration
// @Description Retrieves a specific registration from the database
// @Tags registrations
// @Produce  json
// @Param id path int true "Registration ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Registration
// @Router /registrations/{id} [get]
func (ctrl *RegistrationController) GetRegistration(c *gin.Context) {
	var registration models.Registration
	if result := ctrl.DB.First(&registration, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": registration})
}

// DeleteRegistration godoc
// @Summary Delete a registration
// @Description Deletes a registration from the database
// @Tags registrations
// @Param id path int true "Registration ID"
// @Security ApiKeyAuth
// @Router /registrations/{id} [delete]
func (ctrl *RegistrationController) DeleteRegistration(c *gin.Context) {
	if result := ctrl.DB.Delete(&models.Registration{}, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
