package controllers

import (
	"challenges4/models"
	"challenges4/services"
	"cloud.google.com/go/storage"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func (ctrl *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := ctrl.db.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	ctrl.db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetMe godoc
// @Summary Get the current user
// @Description Get the current user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.PublicUser "Successfully retrieved user"
// @Failure 401 {object} string "Unauthorized: Invalid token"
// @Router /user/me [get]
func (ctrl *UserController) GetMe(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
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
	result := ctrl.db.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateMe godoc
// @Summary Update current user's profile
// @Description Update the profile information of the currently authenticated user, including password.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param first_name formData string false "First name of the user"
// @Param last_name formData string false "Last name of the user"
// @Param email formData string false "Email address of the user"
// @Param profile_picture formData file false "Profile picture file"
// @Param old_password formData string true "Current password for verification"
// @Param new_password formData string false "New password for the user"
// @Success 200 {object} models.PublicUser "Successfully updated user profile"
// @Failure 400 {string} string "Bad request if the provided data is incorrect"
// @Failure 401 {string} string "Unauthorized if the user's old password is incorrect or token is invalid"
// @Failure 404 {string} string "Not Found if the user does not exist"
// @Failure 500 {string} string "Internal Server Error for any server errors"
// @Router /user/me [put]
func (ctrl *UserController) UpdateMe(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
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
	if err := ctrl.db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Retrieve form data
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	email := c.PostForm("email")
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	if file, header, err := c.Request.FormFile("profile_picture"); err == nil {
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close the file", "details": err.Error()})
			}
		}(file)
		if !services.IsValidImageType(header) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Profile picture must be a valid image file"})
			return
		}
		bucketName := os.Getenv("GCS_BUCKET")
		objectName := "users/" + user.Username + "/" + header.Filename
		wc := ctrl.storageService.Client.Bucket(bucketName).Object(objectName).NewWriter(c)
		wc.Metadata = map[string]string{
			"username":   user.Username,
			"uploadTime": time.Now().Format(time.RFC3339),
		}
		defer func(wc *storage.Writer) {
			err := wc.Close()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finalize file upload", "details": err.Error()})
			}
		}(wc)
		if _, err = io.Copy(wc, file); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file", "details": err.Error()})
			return
		}
		url, err := services.GenerateSignedURL(bucketName, objectName, ctrl.storageService.Client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate signed URL", "details": err.Error()})
			return
		}
		user.ProfilePicture = url
	} else if !errors.Is(err, http.ErrMissingFile) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Profile picture is required", "details": err.Error()})
		return
	}

	// Check if the old password is correct before updating
	if oldPassword != "" && newPassword != "" {
		if !services.CheckPasswordHash(oldPassword, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Old password is incorrect"})
			return
		}
		hashedPassword, err := services.HashPassword(newPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash new password", "details": err.Error()})
			return
		}
		user.Password = hashedPassword
	}

	// Update fields
	if firstName != "" {
		user.FirstName = firstName
	}
	if lastName != "" {
		user.LastName = lastName
	}
	if email != "" && email != user.Email {
		var emailCount int64
		ctrl.db.Model(&models.User{}).Where("email = ? AND id <> ?", email, user.ID).Count(&emailCount)
		if emailCount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "This email is already in use by another user"})
			return
		}
		user.Email = email
	}

	// Save changes
	if err := ctrl.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User updated successfully"})
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := ctrl.db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found", "details": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.db.Save(&user)
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.db.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "User deleted"})
}
