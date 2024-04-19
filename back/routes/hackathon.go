package routes

import (
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
// @Security Bearer
// @Success 201 {object} models.Hackathon
// @Router /hackathons [post]
func CreateHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.CreateHackathon(c, db)
	}
}

// GetHackathonsHandler récupère tous les hackathons
// @Summary Lire tous les hackathons
// @Description Récupère une liste de tous les hackathons
// @Tags hackathons
// @Produce  json
// @Security Bearer
// @Success 200 {array} models.Hackathon
// @Router /hackathons [get]
func GetHackathonsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.GetHackathons(c, db)
	}
}

// GetHackathonHandler récupère un hackathon spécifique
// @Summary Lire un hackathon spécifique
// @Description Récupère un hackathon par son ID
// @Tags hackathons
// @Produce  json
// @Param id path int true "ID du Hackathon"
// @Security Bearer
// @Success 200 {object} models.Hackathon
// @Router /hackathons/{id} [get]
func GetHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.GetHackathon(c, db)
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
// @Security Bearer
// @Success 200 {object} models.Hackathon
// @Router /hackathons/{id} [put]
func UpdateHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.UpdateHackathon(c, db)
	}
}

// DeleteHackathonHandler supprime un hackathon
// @Summary Supprimer un hackathon
// @Description Supprime un hackathon par son ID
// @Tags hackathons
// @Produce  json
// @Param id path int true "ID du Hackathon"
// @Security Bearer
// @Success 200 {object} bool "true si la suppression est réussie"
// @Router /hackathons/{id} [delete]
func DeleteHackathonHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.DeleteHackathon(c, db)
	}
}

func HackathonRoutes(r *gin.Engine, db *gorm.DB) {
	hackathonGroup := r.Group("/hackathons")
	hackathonGroup.Use(middleware.AuthMiddleware(0))

	{
		hackathonGroup.POST("/", middleware.AuthMiddleware(2), CreateHackathonHandler(db))
		hackathonGroup.GET("/", GetHackathonsHandler(db))
		hackathonGroup.GET("/:id", GetHackathonHandler(db))
		hackathonGroup.PUT("/:id", middleware.AuthMiddleware(2), UpdateHackathonHandler(db))
		hackathonGroup.DELETE("/:id", middleware.AuthMiddleware(2), DeleteHackathonHandler(db))
	}
}
