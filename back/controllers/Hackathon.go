package controllers

import (
	"challenges4/models"
	"challenges4/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateHackathon(c *gin.Context, db *gorm.DB) {
	var hackathon models.Hackathon

	if err := c.ShouldBindJSON(&hackathon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&hackathon); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

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

	participation := models.Participation{
		HackathonID: hackathon.ID,
		UserID:      user.ID,
		IsOrganizer: true,
	}

	if createParticipationResult := db.Create(&participation); createParticipationResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createParticipationResult.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": hackathon})
}

func GetHackathons(c *gin.Context, db *gorm.DB) {
	var hackathons []models.Hackathon
	if result := db.Find(&hackathons); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hackathons})
}

func GetHackathon(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var hackathon models.Hackathon
	if result := db.First(&hackathon, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hackathon})
}

func UpdateHackathon(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var hackathon models.Hackathon
	if err := db.First(&hackathon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

	var input models.Hackathon
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&hackathon).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": hackathon})
}

func DeleteHackathon(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&models.Hackathon{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func HackathonRegister(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var hackathon models.Hackathon
	if err := db.First(&hackathon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}

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

	var participation models.Participation
	participation.HackathonID = hackathon.ID
	participation.UserID = user.ID

	if result := db.Create(&participation); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Successfully registered for hackathon"})

}

func isHackathonOrganizer(db *gorm.DB, hackathonID uint, userID uint) bool {
	var participation models.Participation
	if result := db.Where("hackathon_id = ? AND user_id = ?", hackathonID, userID).First(&participation); result.Error != nil {
		return false
	}

	return participation.IsOrganizer
}
