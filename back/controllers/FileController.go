package controllers

import (
	"challenges4/models"
	"challenges4/services"
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
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
	userFolder := strconv.Itoa(int(userID))
	objectName := userFolder + "/" + header.Filename

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

	url, err := services.GenerateSignedURL(bucketName, objectName, storageService.Client)

	c.JSON(http.StatusOK, gin.H{"url": url})
}

// GetMyFiles godoc
// @Summary Get all files uploaded by the user
// @Description Get all files uploaded by the user from Google Cloud Storage
// @Tags files
// @Security ApiKeyAuth
// @Success 200 {object} string "List of files"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /files/me [get]
func GetMyFiles(c *gin.Context, storageService *services.StorageService) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
		return
	}

	if len(token) > 7 && strings.ToUpper(token[:7]) == "BEARER " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx := context.Background()
	bucketName := os.Getenv("GCS_BUCKET")
	userFolder := strconv.Itoa(int(userID))

	query := &storage.Query{Prefix: userFolder + "/"}

	it := storageService.Client.Bucket(bucketName).Objects(ctx, query)

	var fileInfos []models.FileInfo
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list files: " + err.Error()})
			return
		}
		// Filter files by userID
		if attrs.Metadata["userID"] == strconv.Itoa(int(userID)) {
			url, err := services.GenerateSignedURL(bucketName, attrs.Name, storageService.Client)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate signed URL: " + err.Error()})
				return
			}
			splitPath := strings.Split(attrs.Name, "/")
			filename := splitPath[len(splitPath)-1]
			fileInfos = append(fileInfos, models.FileInfo{
				Name:       filename,
				UploadTime: attrs.Metadata["uploadTime"],
				URL:        url,
			})
		}
	}

	if fileInfos == nil {
		fileInfos = []models.FileInfo{}
	}
	c.JSON(http.StatusOK, gin.H{"files": fileInfos})
}
