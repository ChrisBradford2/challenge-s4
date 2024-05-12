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
}
