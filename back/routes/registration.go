package routes

import (
	"challenges4/config"
	"challenges4/controllers"
	middlewares "challenges4/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRegistrationRoutes(router *gin.Engine, db *gorm.DB) {
	registrationController := controllers.NewRegistrationController(db)

	// CreateRegistrationHandler creates a registration (accessible by RoleUser)
	// @Summary Create a registration
	// @Description Adds a new registration to the database
	// @Tags registrations
	// @Accept  json
	// @Produce  json
	// @Param registration body models.RegistrationCreate true "Registration to create"
	// @Security ApiKeyAuth
	// @Success 201 {object} models.Registration
	// @Router /registrations [post]
	router.POST("/registrations", middlewares.AuthMiddleware(config.RoleUser), func(c *gin.Context) {
		registrationController.CreateRegistration(c)
	})

	// GetRegistrationsHandler retrieves all registrations
	// @Summary Get all registrations
	// @Description Retrieves all registrations from the database
	// @Tags registrations
	// @Produce  json
	// @Security ApiKeyAuth
	// @Success 200 {array} models.Registration
	// @Router /registrations [get]
	router.GET("/registrations", middlewares.AuthMiddleware(config.RoleUser), func(c *gin.Context) {
		registrationController.GetRegistrations(c)
	})

	// GetRegistrationHandler retrieves a specific registration
	// @Summary Get a registration
	// @Description Retrieves a specific registration from the database
	// @Tags registrations
	// @Produce  json
	// @Param id path int true "Registration ID"
	// @Security ApiKeyAuth
	// @Success 200 {object} models.Registration
	// @Router /registrations/{id} [get]
	router.GET("/registrations/:id", middlewares.AuthMiddleware(config.RoleUser), func(c *gin.Context) {
		registrationController.GetRegistration(c)
	})

	// UpdateRegistrationHandler updates a registration
	// @Summary Update a registration
	// @Description Updates a registration in the database
	// @Tags registrations
	// @Accept  json
	// @Produce  json
	// @Param id path int true "Registration ID"
	// @Param registration body models.RegistrationUpdate true "Registration to update"
	// @Security ApiKeyAuth
	// @Success 200 {object} models.Registration
	// @Router /registrations/{id} [patch]
	router.PATCH("/registrations/:id", middlewares.AuthMiddleware(config.RoleUser), func(c *gin.Context) {
		registrationController.UpdateRegistration(c)
	})
}
