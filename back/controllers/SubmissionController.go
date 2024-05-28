package controllers

//
//import (
//	"challenges4/models"
//	"github.com/gin-gonic/gin"
//	"gorm.io/gorm"
//	"net/http"
//)
//
//type SubmissionController struct {
//	DB *gorm.DB
//}
//
//func NewSubmissionController(db *gorm.DB) *SubmissionController {
//	return &SubmissionController{DB: db}
//}
//
//// CreateSubmission godoc
//// @Summary Create a new Submission
//// @Description Create a new Submission
//// @Tags Submissions
//// @Accept  json
//// @Produce  json
//// @Security ApiKeyAuth
//// @Success 200 {object} models.Submission "Successfully created Submission
//// @Failure 400 {object} string "Bad request"
//// @Failure 500 {object} string "Internal server error"
//// @Router /submissions [post]
//func (ctrl *SubmissionController) CreateSubmission(c *gin.Context) {
//	var submission models.Submission
//	//log the request
//	// Log the request details
//	//log.Printf("Request Body: %s", string(requestBody))
//
//	if err := c.ShouldBindJSON(&submission); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	if result := ctrl.DB.Create(&submission); result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{"data": submission})
//}
//
//// UpdateSubmission godoc
//// @Summary Update a Submission
//// @Description Update a Submission by ID
//// @Tags Submissions
//// @Accept  json
//// @Produce  json
//// @Param id path int true "Submission ID"
//// @Param Submission body models.Submission true "Submission object"
//// @Security ApiKeyAuth
//// @Success 200 {object} models.Submission "Successfully updated Submission"
//// @Failure 400 {object} string "Bad request"
//// @Failure 404 {object} string "Submission not found"
//func (ctrl *SubmissionController) UpdateSubmission(c *gin.Context) {
//	id := c.Param("id")
//	var submission models.Submission
//	if err := ctrl.DB.First(&submission, id).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
//		return
//	}
//
//	var input models.Submission
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	submission.GitLink = input.Title
//	submission.Description = input.Description
//	submission.HackathonID = input.HackathonID
//	submission.UserID = input.UserID
//
//	if result := ctrl.DB.Save(&submission); result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"data": submission})
//}
