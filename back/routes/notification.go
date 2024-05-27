package routes

import (
	"challenges4/config"
	"challenges4/controllers"
	middleware "challenges4/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateNotificationHandler crée une notification
// @Summary Créer une notification
// @Description Ajoute une nouvelle notification à la base de données
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param notification body models.NotificationCreate true "Notification à créer"
// @Security ApiKeyAuth
// @Success 201 {object} models.Notification
// @Router /notifications [post]
func CreateNotificationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.CreateNotification(c, db)
	}
}

// GetNotificationsHandler récupère toutes les notifications
// @Summary Lire toutes les notifications
// @Description Récupère une liste de toutes les notifications
// @Tags notifications
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.Notification
// @Router /notifications [get]
func GetNotificationsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.GetNotifications(c, db)
	}
}

// GetNotificationHandler récupère une notification spécifique
// @Summary Lire une notification spécifique
// @Description Récupère une notification par son ID
// @Tags notifications
// @Produce  json
// @Param id path int true "ID de la Notification"
// @Security ApiKeyAuth
// @Success 200 {object} models.Notification
// @Router /notifications/{id} [get]
func GetNotificationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.GetNotification(c, db)
	}
}

// UpdateNotificationHandler met à jour une notification
// @Summary Mettre à jour une notification
// @Description Met à jour les informations d'une notification par son ID
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param id path int true "ID de la Notification"
// @Param notification body models.Notification true "Informations de la Notification à mettre à jour"
// @Security ApiKeyAuth
// @Success 200 {object} models.Notification
// @Router /notifications/{id} [put]
func UpdateNotificationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.UpdateNotification(c, db)
	}
}

// DeleteNotificationHandler supprime une notification
// @Summary Supprimer une notification
// @Description Supprime une notification par son ID
// @Tags notifications
// @Produce  json
// @Param id path int true "ID de la Notification"
// @Security ApiKeyAuth
// @Success 200 {object} bool "true si la suppression est réussie"
// @Router /notifications/{id} [delete]
func DeleteNotificationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.DeleteNotification(c, db)
	}
}

func NotificationRoutes(r *gin.Engine, db *gorm.DB) {
	notificationGroup := r.Group("/notifications")
	notificationGroup.Use(middleware.AuthMiddleware(0))

	{
		notificationGroup.POST("/", middleware.AuthMiddleware(config.RoleUser), CreateNotificationHandler(db))
		notificationGroup.GET("/", middleware.AuthMiddleware(config.RoleUser), GetNotificationsHandler(db))
		notificationGroup.GET("/:id", middleware.AuthMiddleware(config.RoleUser), GetNotificationHandler(db))
		notificationGroup.PUT("/:id", middleware.AuthMiddleware(config.RoleUser), UpdateNotificationHandler(db))
		notificationGroup.DELETE("/:id", middleware.AuthMiddleware(config.RoleUser), DeleteNotificationHandler(db))
	}
}
