package routes

import (
	"challenges4/config"
	"challenges4/controllers"
	middleware "challenges4/middlewares"
	"challenges4/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB, storageService *services.StorageService) {
	userGroup := r.Group("/user")
	{
		userController := controllers.NewUserController(db, storageService)
		userGroup.POST("/login", userController.Login)
		userGroup.POST("/register", userController.Register)
		userGroup.GET("/me",
			middleware.AuthMiddleware(config.RoleUser),
			userController.GetMe,
		)
		userGroup.PUT("/me",
			middleware.AuthMiddleware(config.RoleUser),
			userController.UpdateMe,
		)
		userGroup.DELETE("/me",
			middleware.AuthMiddleware(config.RoleUser),
			userController.DeleteMe,
		)
	}
}
