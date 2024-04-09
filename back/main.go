package main

import (
	"challenges4/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/madkins23/gin-utils/pkg/ginzero"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	r := gin.New()
	r.Use(ginzero.Logger())

	//connexion bd
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	// Récupérer les informations de connexion à la base de données à partir des variables d'environnement
	dsn := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"

	// Ouvrir une connexion à la base de données
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connexion à la base de données réussie !")

	// Automatiser la migration de la base de données
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Erreur lors de la migration de la base de données:", err)
	}

	log.Println("Migration de la base de données réussie !")
	// Mot de passe à hacher
	password := "test1234"

	// Générer un hachage bcrypt pour le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erreur lors du hachage du mot de passe:", err)
		return
	}
	nouvelUtilisateur := models.User{Nom: "Dupont", Prenom: "Alice", Login: "aliced", Password: string(hashedPassword)}
	db.Create(&nouvelUtilisateur)
	//connexion bd

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, gin-zerolog example")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":8080")
}
