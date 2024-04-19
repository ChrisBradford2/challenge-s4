package main

import (
	"challenges4/docs"
	"challenges4/models"
	"challenges4/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/madkins23/gin-utils/pkg/ginzero"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func connectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	// Try to connect to the database 5 times
	for i := 0; i < 5; i++ {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		log.Printf("Failed to connect to the database. Retrying in 5 seconds... (attempt %d/5)", i+1)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to the database after 5 attempts")
}

// @title Kiwi Collective API
// @description Swagger API for the Kiwi Collective project.
// @securityDefinitions.api_key Bearer
// @version 1.0
// @BasePath /
func main() {
	docs.SwaggerInfo.Title = "Kiwi Collective API"
	docs.SwaggerInfo.Description = "API for the Kiwi Collective project."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.New()
	r.Use(ginzero.Logger())

	//DB connection
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := connectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Connected to the database !")

	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate the database: ", err)
	}

	log.Println("Database migrated !")

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, gin-zerolog example")
	})

	routes.UserRoutes(r, db)

	routes.HackathonRoutes(r, db)
	if err := db.AutoMigrate(&models.Hackathon{}); err != nil {
		log.Fatal("Failed to migrate the database: ", err)
	}

	// File upload
	routes.FileRoutes(r, os.Getenv("GCP_CREDS"))

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
