package routes

import (
	"challenges4/config"
	"challenges4/controllers"
	middlewares "challenges4/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTeamRoutes(router *gin.Engine, db *gorm.DB) {
	teamController := controllers.NewTeamController(db)

	teamRoutes := router.Group("/teams")
	teamRoutes.Use(middlewares.AuthMiddleware(config.RoleOrganizer | config.RoleAdmin))

	// CreateTeamHandler creates a team
	// @Summary Create a team
	// @Description Adds a new team to the database
	// @Tags teams
	// @Accept  json
	// @Produce  json
	// @Param team body models.TeamCreate true "Team to create"
	// @Security ApiKeyAuth
	// @Success 201 {object} models.Team
	// @Router /teams [post]
	router.POST("/teams", func(c *gin.Context) {
		teamController.CreateTeam(c)
	})

	// GetTeamsHandler retrieves all teams
	// @Summary Get all teams
	// @Description Retrieves all teams from the database
	// @Tags teams
	// @Produce  json
	// @Security ApiKeyAuth
	// @Success 200 {array} models.Team
	// @Router /teams [get]
	router.GET("/teams", func(c *gin.Context) {
		teamController.GetTeams(c)
	})

	// GetTeamHandler retrieves a specific team
	// @Summary Get a team
	// @Description Retrieves a specific team from the database
	// @Tags teams
	// @Produce  json
	// @Param id path int true "Team ID"
	// @Security ApiKeyAuth
	// @Success 200 {object} models.Team
	// @Router /teams/{id} [get]
	router.GET("/teams/:id", func(c *gin.Context) {

		teamController.GetTeam(c)
	})

	// UpdateTeamHandler updates a team
	// @Summary Update a team
	// @Description Updates a team in the database
	// @Tags teams
	// @Accept  json
	// @Produce  json
	// @Param id path int true "Team ID"
	// @Param team body models.TeamUpdate true "Team to update"
	// @Security ApiKeyAuth
	// @Success 200 {object} models.Team
	// @Router /teams/{id} [put]
	router.PUT("/teams/:id", func(c *gin.Context) {
		teamController.UpdateTeam(c)
	})

	// DeleteTeamHandler deletes a team
	// @Summary Delete a team
	// @Description Deletes a team from the database
	// @Tags teams
	// @Param id path int true "Team ID"
	// @Security ApiKeyAuth
	// @Success 204
	// @Router /teams/{id} [delete]
	router.DELETE("/teams/:id", func(c *gin.Context) {
		teamController.DeleteTeam(c)
	})

	// RegisterToTeamHandler registers a user to a team
	// @Summary Register to a team
	// @Description Register to a team by ID
	// @Tags teams
	// @Produce  json
	// @Param id path int true "Team ID"
	// @Security ApiKeyAuth
	// @Success 200 {object} models.Team "Successfully registered to team"
	// @Failure 400 {object} string "User is already in a team or Team is already full"
	// @Failure 401 {object} string "Unauthorized: No authorization token provided or invalid token"
	// @Failure 404 {object} string "Team not found or User not found"
	// @Failure 500 {object} string "Internal server error"
	// @Router /teams/{id}/register [post]
	router.POST("/teams/:id/register", func(c *gin.Context) {
		teamController.RegisterToTeam(c)
	})

	// LeaveTeamHandler removes a user from a team
	// @Summary Leave a team
	// @Description Leave a team by ID
	// @Tags teams
	// @Produce  json
	// @Param id path int true "Team ID"
	// @Security ApiKeyAuth
	// @Success 200 {object} string "Successfully left team"
	// @Failure 400 {object} string "User is not in the target team"
	// @Failure 401 {object} string "Unauthorized: No authorization token provided or invalid token"
	// @Failure 404 {object} string "Team not found or User not found"
	// @Failure 500 {object} string "Internal server error"
	// @Router /teams/{id}/leave [post]
	router.POST("/teams/:id/leave", func(c *gin.Context) {
		teamController.LeaveTeam(c)
	})
}
