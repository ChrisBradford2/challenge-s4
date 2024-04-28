package controllers

import (
	"challenges4/models"
	"challenges4/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
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
// @Accept json
// @Produce json
// @Param user body models.UserRegister true "User object"
// @Success 201 {object} models.UserRegisterResponse "User registered"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /user/register [post]
func (ctrl *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	hashedPassword, err := services.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = hashedPassword

	if user.TeamID != nil {
		exists := ctrl.db.First(&models.Team{}, *user.TeamID).Error == nil
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Team not found"})
			return
		}
	}

	if err := ctrl.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	response := models.UserRegisterResponse{
		Username: user.Username,
		Email:    user.Email,
	}
	c.JSON(http.StatusCreated, response)
}
