package routes

import (
	"challenges4/config"
	"challenges4/controllers"
	middleware "challenges4/middlewares"
	"challenges4/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SubmissionRoutes(r *gin.Engine, db *gorm.DB, storageService *services.StorageService) {
	submissionController := controllers.NewSubmissionController(db)

	submissionGroup := r.Group("/submissions")
	{
		submissionGroup.Use(middleware.AuthMiddleware(config.RoleUser))
		submissionGroup.POST("/", submissionController.CreateSubmission)
		submissionGroup.POST("/upload", func(c *gin.Context) {
			submissionController.UploadSubmissionFile(c, storageService)
		})
	}
}
