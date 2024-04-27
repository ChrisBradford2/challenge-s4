package routes

import (
	"challenges4/controllers"
	middleware "challenges4/middlewares"
	"challenges4/services"
	"context"
	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.Engine, credentialsFile string) {
	storageService := services.NewStorageService(context.Background(), credentialsFile)

	r.POST("/upload",
		middleware.AuthMiddleware(0),
		func(c *gin.Context) {
			controllers.UploadFile(c, storageService)
		})
	r.GET("/files/me", middleware.AuthMiddleware(0), func(c *gin.Context) {
		controllers.GetMyFiles(c, storageService)
	})
}
