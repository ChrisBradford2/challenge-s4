package controllers

import (
	"challenges4/config"
	"challenges4/models"
	"challenges4/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type StepController struct {
	DB *gorm.DB
}

type JSONResponse map[string]interface{}

func NewStepController(db *gorm.DB) *StepController {
	return &StepController{DB: db}
}

// CreateStep godoc
// @Summary Create a new step
// @Description Create a new step for a hackathon
// @Tags steps
// @Accept  json
// @Produce  json
// @Param step body models.StepCreate true "Step object"
// @Security ApiKeyAuth
// @Success 201 {object} JSONResponse{message=string,data=models.Step} "Step created successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Internal server error"
// @Router /steps [post]
func (ctrl *StepController) CreateStep(c *gin.Context) {
	var step models.StepCreate

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
	if err := ctrl.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&step); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	if user.Roles != config.RoleAdmin && !isHackathonOrganizer(ctrl.DB, *step.HackathonID, user.ID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User is not an admin or organizer of the hackathon"})
		return
	}

	if result := ctrl.DB.Create(&step); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, JSONResponse{
		"message": "Step created successfully",
		"data":    step,
	})
}

// UpdateStep godoc
// @Summary Update a step
// @Description Update a step by ID
// @Tags steps
// @Accept  json
// @Produce  json
// @Param id path int true "Step ID"
// @Param step body models.StepUpdate true "Step object"
// @Security ApiKeyAuth
// @Success 200 {object} JSONResponse{message=string,data=models.Step} "Step updated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Step not found"
// @Failure 500 {object} string "Internal server error"
// @Router /steps/{id} [put]
func (ctrl *StepController) UpdateStep(c *gin.Context) {
	id := c.Param("id")
	var step models.StepUpdate

	var existingStep models.Step
	if err := ctrl.DB.First(&existingStep, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Step not found"})
		return
	}

	if err := c.ShouldBindJSON(&step); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
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
	if err := ctrl.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Roles != config.RoleAdmin && !isHackathonOrganizer(ctrl.DB, *existingStep.HackathonID, user.ID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User is not an admin or organizer of the hackathon"})
		return
	}

	if result := ctrl.DB.Save(&step); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Step updated successfully",
		"data":    step,
	})
}
