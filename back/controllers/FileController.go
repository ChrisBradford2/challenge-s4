package controllers

import (
	"challenges4/services"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// UploadFile godoc
// @Summary Upload a file
// @Description Uploads a file to Google Cloud Storage
// @Tags files
// @Accept multipart/form-data
// @Param file formData file true "File to upload"
// @Security ApiKeyAuth
// @Success 200 {string} string "File uploaded successfully"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /upload [post]
func UploadFile(c *gin.Context, storageService *services.StorageService) {
	token := c.GetHeader("Authorization") // Get token from header
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
		return
	}

	// Delete "Bearer " from the token
	if len(token) > 7 && strings.ToUpper(token[:7]) == "BEARER " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}(file)

	ctx := context.Background()
	bucketName := os.Getenv("GCS_BUCKET")
	objectName := header.Filename

	wc := storageService.Client.Bucket(bucketName).Object(objectName).NewWriter(ctx)

	wc.Metadata = map[string]string{
		"userID":     strconv.Itoa(int(userID)),
		"uploadTime": time.Now().Format(time.RFC3339),
	}

	if _, err = io.Copy(wc, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := wc.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
