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
	"time"
)

type UserController struct {
	db             *gorm.DB
	storageService *services.StorageService
}

func NewUserController(db *gorm.DB, storageService *services.StorageService) *UserController {
	return &UserController{
		db:             db,
		storageService: storageService,
	}
}

// Login godoc
// @Summary Login
// @Description Logs in a user
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.UserLogin true "User credentials"
// @Success 200 {string} string "Token"
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Failure 401 {object} models.ErrorResponse "Invalid request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /user/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var credentials models.UserLogin
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	user, err := models.GetUserByEmail(ctrl.db, credentials.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	if !services.CheckPasswordHash(credentials.Password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := services.GenerateJWT(*user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

// Register godoc
// @Summary Register
// @Description Registers a new user
// @Tags auth
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Username"
// @Param last_name formData string true "Last Name"
// @Param first_name formData string true "First Name"
// @Param profile_picture formData file true "Profile Picture"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 201 {object} models.UserRegisterResponse "User registered"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /user/register [post]
func (ctrl *UserController) Register(c *gin.Context) {
	// Extract text fields
	username := c.PostForm("username")
	lastName := c.PostForm("last_name")
	firstName := c.PostForm("first_name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Extract file
	file, header, err := c.Request.FormFile("profile_picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Profile picture is required", "details": err.Error()})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close the file", "details": err.Error()})
		}
	}(file)

	// Check if the file is an image
	if !services.IsValidImageType(header) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
		return
	}

	// Check if the email is already in use
	var count int64
	ctrl.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Check if the user already exists
	ctrl.db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Upload to Google Cloud Storage and get the URL
	bucketName := os.Getenv("GCS_BUCKET")
	objectName := "users/" + username + "/" + header.Filename

	wc := ctrl.storageService.Client.Bucket(bucketName).Object(objectName).NewWriter(context.Background())
	wc.Metadata = map[string]string{
		"username":   username,
		"uploadTime": time.Now().Format(time.RFC3339),
	}
	if _, err = io.Copy(wc, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not upload profile picture"})
		return
	}
	if err := wc.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	url, err := services.GenerateSignedURL(bucketName, objectName, ctrl.storageService.Client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate file URL"})
		return
	}

	// Hash password
	hashedPassword, err := services.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Create user model
	user := models.User{
		Username:       username,
		LastName:       lastName,
		FirstName:      firstName,
		ProfilePicture: url,
		Email:          email,
		Password:       hashedPassword,
	}

	// Insert user into database
	if err := ctrl.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// Prepare response
	response := models.UserRegisterResponse{
		Username: username,
		Email:    email,
	}
	c.JSON(http.StatusCreated, response)
}
