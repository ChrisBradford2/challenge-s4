package routes

import (
	"challenges4/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/user")
	{
		userController := controllers.NewUserController(db)
		userGroup.POST("/login", userController.Login)
		userGroup.POST("/register", userController.Register)
	}
}
