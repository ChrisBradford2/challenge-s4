package routes

import (
	"challenges4/config"
	"challenges4/controllers"
	middleware "challenges4/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.CreateHackathon(c, db)
	}
}

func GetHackathonsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.GetHackathons(c, db)
	}
}

func GetHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.GetHackathon(c, db)
	}
}

func UpdateHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.UpdateHackathon(c, db)
	}
}

func DeleteHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.DeleteHackathon(c, db)
	}
}

func RegisterHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.HackathonRegister(c, db)
	}
}

func HackathonRoutes(r *gin.Engine, db *gorm.DB) {
	hackathonGroup := r.Group("/hackathons")
	hackathonGroup.Use(middleware.AuthMiddleware(0))

	{
		hackathonGroup.POST("/", middleware.AuthMiddleware(config.RoleUser), CreateHackathonHandler(db))
		hackathonGroup.GET("/", middleware.AuthMiddleware(config.RoleUser), GetHackathonsHandler(db))
		hackathonGroup.GET("/:id", middleware.AuthMiddleware(config.RoleUser), GetHackathonHandler(db))
		hackathonGroup.PUT("/:id", middleware.AuthMiddleware(config.RoleOrganizer), UpdateHackathonHandler(db))
		hackathonGroup.DELETE("/:id", middleware.AuthMiddleware(config.RoleOrganizer), DeleteHackathonHandler(db))
		hackathonGroup.POST("/:id/register", middleware.AuthMiddleware(config.RoleUser), RegisterHackathonHandler(db))
	}
}
