package controllers

import (
	"challenges4/models"
	"challenges4/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// CreateNotification creates a new notification
func CreateNotification(c *gin.Context, db *gorm.DB) {
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
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	/*if !user.isOrganizer { // utilisateur lambda

	}*/

	var notification models.Notification

	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&notification); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": notification})
}

// GetNotifications récupérer toutes les notifications
func GetNotifications(c *gin.Context, db *gorm.DB) {
	var notifications []models.Notification
	if result := db.Find(&notifications); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notifications})
}

// GetNotification retrouver une notification pas id
func GetNotification(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var notification models.Notification
	if result := db.First(&notification, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notification})
}

// UpdateNotification modifier une notification
func UpdateNotification(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var notification models.Notification
	if err := db.First(&notification, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	var input models.Notification
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&notification).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": notification})
}

// DeleteNotification supprimer une notification par id
func DeleteNotification(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&models.Notification{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
