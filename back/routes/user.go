package routes

import (
	"challenges4/controllers"
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
	}
}
