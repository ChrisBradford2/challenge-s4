package controllers

import (
	"challenges4/models"
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
	var registration models.Registration
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
// @Param registration body models.Registration true "Registration object"
// @Security ApiKeyAuth
// @Success 200 {object} models.Registration "Successfully updated registration"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Registration not found"
// @Failure 500 {object} string "Internal server error"
// @Router /registrations/{id} [patch]
func (ctrl *RegistrationController) UpdateRegistration(c *gin.Context) {
	var registration models.Registration
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := ctrl.DB.Save(&registration); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": registration})
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
	query := ctrl.DB

	// Get query parameters
	teamID := c.Query("teamId")
	user := c.Query("user")
	status := c.Query("status")

	// Apply filters based on query parameters
	if teamID != "" {
		tid, err := strconv.Atoi(teamID)
		if err == nil {
			query = query.Where("team_id = ?", tid)
		}
	}
	if user != "" {
		query = query.Where("user = ?", user)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if result := query.Find(&registrations); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": registrations})
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
