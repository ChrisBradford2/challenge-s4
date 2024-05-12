package controllers

import (
	"challenges4/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type TeamController struct {
	DB *gorm.DB
}

func NewTeamController(db *gorm.DB) *TeamController {
	return &TeamController{DB: db}
}

// CreateTeam godoc
// @Summary Create a new team
// @Description Create a new team
// @Tags teams
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Team "Successfully retrieved team"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /teams [post]
func (ctrl *TeamController) CreateTeam(c *gin.Context) {
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil { // Check if the request body is valid JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := ctrl.DB.Create(&team); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": team})
}

// UpdateTeam godoc
// @Summary Update a team
// @Description Update a team by ID
// @Tags teams
// @Accept  json
// @Produce  json
// @Param id path int true "Team ID"
// @Param team body models.Team true "Team object"
// @Security ApiKeyAuth
// @Success 200 {object} models.Team "Successfully updated team"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Team not found"
// @Failure 500 {object} string "Internal server error"
// @Router /teams/{id} [patch]
func (ctrl *TeamController) UpdateTeam(c *gin.Context) {
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := ctrl.DB.Model(&models.Team{}).Where("id = ?", c.Param("id")).Updates(&team); result.Error != nil {
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": team})
}

// GetTeams godoc
// @Summary Get all teams
// @Description Get all teams
// @Tags teams
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.Team "Successfully retrieved teams"
// @Failure 500 {object} string "Internal server error"
// @Router /teams [get]
func (ctrl *TeamController) GetTeams(c *gin.Context) {
	var teams []models.Team
	if result := ctrl.DB.Find(&teams); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teams})
}

// GetTeam godoc
// @Summary Get a team
// @Description Get a team by ID
// @Tags teams
// @Produce  json
// @Param id path int true "Team ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Team "Successfully retrieved team"
// @Failure 404 {object} string "Team not found"
// @Failure 500 {object} string "Internal server error"
// @Router /teams/{id} [get]
func (ctrl *TeamController) GetTeam(c *gin.Context) {
	var team models.Team
	if result := ctrl.DB.First(&team, c.Param("id")); result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": team})
}

// DeleteTeam godoc
// @Summary Delete a team
// @Description Delete a team by ID
// @Tags teams
// @Produce  json
// @Param id path int true "Team ID"
// @Security ApiKeyAuth
// @Success 204 "Successfully deleted team"
// @Failure 404 {object} string "Team not found"
// @Failure 500 {object} string "Internal server error"
// @Router /teams/{id} [delete]
func (ctrl *TeamController) DeleteTeam(c *gin.Context) {
	if result := ctrl.DB.Delete(&models.Team{}, c.Param("id")); result.Error != nil {
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
