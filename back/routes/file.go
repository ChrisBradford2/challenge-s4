package routes

import (
	"challenges4/controllers"
	"challenges4/services"
	"context"
	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.Engine, credentialsFile string) {
	storageService := services.NewStorageService(context.Background(), credentialsFile)

	r.POST("/upload", func(c *gin.Context) {
		controllers.UploadFile(c, storageService)
	})
}
