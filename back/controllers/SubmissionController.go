package controllers

import (
	"challenges4/models"
	"challenges4/services"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type SubmissionController struct {
	DB *gorm.DB
}

func NewSubmissionController(db *gorm.DB) *SubmissionController {
	return &SubmissionController{DB: db}
}

// CreateSubmission godoc
// @Summary Create a new submission
// @Description Create a new submission for a step
// @Tags submissions
// @Accept  json
// @Produce  json
// @Param submission body models.SubmissionCreate true "Submission object to create"
// @Security ApiKeyAuth
// @Success 201 {object} JSONResponse{message=string,data=models.Submission} "Submission created successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /submissions [post]
func (ctrl *SubmissionController) CreateSubmission(c *gin.Context) {
	var submission models.Submission

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	var user models.User
	if err := ctrl.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return
	}

	submission.Status = "submitted"
	if result := ctrl.DB.Create(&submission); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, JSONResponse{
		"message": "Submission created successfully",
		"data":    submission,
	})
}

// UploadSubmissionFile godoc
// @Summary Upload a file for a submission
// @Description Uploads a file to Google Cloud Storage for a submission
// @Tags submissions
// @Accept multipart/form-data
// @Param file formData file true "File to upload"
// @Security ApiKeyAuth
// @Success 201 {object} JSONResponse{message=string,url=string} "File uploaded successfully"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /submissions/upload [post]
func (ctrl *SubmissionController) UploadSubmissionFile(c *gin.Context, storageService *services.StorageService) {
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "details": err.Error()})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file", "details": err.Error()})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close the file", "details": err.Error()})
		}
	}(file)

	ctx := context.Background()
	bucketName := os.Getenv("GCS_BUCKET")
	objectName := "submissions/" + strconv.Itoa(int(userID)) + "/" + header.Filename

	wc := storageService.Client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	wc.Metadata = map[string]string{
		"userID":     strconv.Itoa(int(userID)),
		"uploadTime": time.Now().Format(time.RFC3339),
	}

	if _, err = io.Copy(wc, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file", "details": err.Error()})
		return
	}
	if err := wc.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finalize file upload", "details": err.Error()})
		return
	}

	url, err := services.GenerateSignedURL(bucketName, objectName, storageService.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate signed URL", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, JSONResponse{
		"message": "File uploaded successfully",
		"url":     url,
	})
}
