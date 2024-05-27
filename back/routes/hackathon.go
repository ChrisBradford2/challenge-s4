package routes

import (
	"challenges4/config"
	"challenges4/controllers"
	middleware "challenges4/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateHackathonHandler crée un hackathon
// @Summary Créer un hackathon
// @Description Ajoute un nouveau hackathon à la base de données
// @Tags hackathons
// @Accept  json
// @Produce  json
// @Param hackathon body models.HackathonCreate true "Hackathon à créer"
// @Security ApiKeyAuth
// @Success 201 {object} models.Hackathon
// @Router /hackathons [post]
func CreateHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.CreateHackathon(c, db)
	}
}

// GetHackathonsHandler récupère tous les hackathons
// @Summary Lire tous les hackathons
// @Description Récupère une liste de tous les hackathons
// @Tags hackathons
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.Hackathon
// @Router /hackathons [get]
func GetHackathonsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.GetHackathons(c, db)
	}
}

// GetHackathonHandler récupère un hackathon spécifique
// @Summary Lire un hackathon spécifique
// @Description Récupère un hackathon par son ID
// @Tags hackathons
// @Produce  json
// @Param id path int true "ID du Hackathon"
// @Security ApiKeyAuth
// @Success 200 {object} models.Hackathon
// @Router /hackathons/{id} [get]
func GetHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.GetHackathon(c, db)
	}
}

// UpdateHackathonHandler met à jour un hackathon
// @Summary Mettre à jour un hackathon
// @Description Met à jour les informations d'un hackathon par son ID
// @Tags hackathons
// @Accept  json
// @Produce  json
// @Param id path int true "ID du Hackathon"
// @Param hackathon body models.Hackathon true "Informations du Hackathon à mettre à jour"
// @Security ApiKeyAuth
// @Success 200 {object} models.Hackathon
// @Router /hackathons/{id} [put]
func UpdateHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.UpdateHackathon(c, db)
	}
}

// DeleteHackathonHandler supprime un hackathon
// @Summary Supprimer un hackathon
// @Description Supprime un hackathon par son ID
// @Tags hackathons
// @Produce  json
// @Param id path int true "ID du Hackathon"
// @Security ApiKeyAuth
// @Success 200 {object} bool "true si la suppression est réussie"
// @Router /hackathons/{id} [delete]
func DeleteHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.DeleteHackathon(c, db)
	}
}

// RegisterHackathonHandler enregistre un utilisateur à un hackathon
// @Summary Enregistrer un utilisateur à un hackathon
// @Description Enregistre un utilisateur à un hackathon par son ID
// @Tags hackathons
// @Produce  json
// @Param id path int true "ID du Hackathon"
// @Security ApiKeyAuth
// @Success 200 {object} bool "Inscription réussie"
// @Router /hackathons/{id}/register [post]
func RegisterHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.HackathonRegister(c, db)
	}
}

func HackathonRoutes(r *gin.Engine, db *gorm.DB) {
	hackathonGroup := r.Group("/hackathons")
	hackathonGroup.Use(middleware.AuthMiddleware(0))

	{
		hackathonGroup.POST("/", middleware.AuthMiddleware(config.RoleOrganizer), CreateHackathonHandler(db))
		hackathonGroup.GET("/", middleware.AuthMiddleware(config.RoleUser), GetHackathonsHandler(db))
		hackathonGroup.GET("/:id", middleware.AuthMiddleware(config.RoleUser), GetHackathonHandler(db))
		hackathonGroup.PUT("/:id", middleware.AuthMiddleware(config.RoleOrganizer), UpdateHackathonHandler(db))
		hackathonGroup.DELETE("/:id", middleware.AuthMiddleware(config.RoleOrganizer), DeleteHackathonHandler(db))
		hackathonGroup.POST("/:id/register", middleware.AuthMiddleware(config.RoleUser), RegisterHackathonHandler(db))
	}
}
