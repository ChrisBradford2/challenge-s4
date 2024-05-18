package controllers

import (
	"challenges4/models"
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
