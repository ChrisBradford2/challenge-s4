package models

import "gorm.io/gorm"

type Utilisateur struct {
	gorm.Model
	ID       uint `gorm:"not null"`
	Nom      string
	Prenom   string
	Login    string `gorm:"not null"`
	Email    string
	Password string `gorm:"not null"`
	// Ajoutez d'autres champs selon votre structure de base de données
}
