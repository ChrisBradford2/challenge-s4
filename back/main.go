package main

import (
	"challenges4/database"
	"challenges4/docs"
	"challenges4/models"
	"challenges4/routes"
	"challenges4/seeders"
	"challenges4/services"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/madkins23/gin-utils/pkg/ginzero"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	"time"
)

// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
//
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer token
//
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Database connection
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Connected to the database!")

	// Migrate the schema in a controlled order with detailed logging
	log.Println("Starting migrations...")

	log.Println("Migrating teams...")
	if err := db.AutoMigrate(&models.Team{}); err != nil {
		log.Fatal("Failed to migrate teams: ", err)
	}
	log.Println("Migrated teams!")

	log.Println("Migrating users...")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate users: ", err)
	}
	log.Println("Migrated users!")

	log.Println("Migrating hackathons...")
	if err := db.AutoMigrate(&models.Hackathon{}); err != nil {
		log.Fatal("Failed to migrate hackathons: ", err)
	}
	log.Println("Migrated hackathons!")

	log.Println("Migrating files...")
	if err := db.AutoMigrate(&models.File{}); err != nil {
		log.Fatal("Failed to migrate files: ", err)
	}
	log.Println("Migrated files!")

	log.Println("Migrating skills...")
	if err := db.AutoMigrate(&models.Skill{}); err != nil {
		log.Fatal("Failed to migrate skills: ", err)
	}
	log.Println("Migrated skills!")

	log.Println("Migrating steps...")
	if err := db.AutoMigrate(&models.Step{}); err != nil {
		log.Fatal("Failed to migrate steps: ", err)
	}
	log.Println("Migrated steps!")

	log.Println("Migrating participations...")
	if err := db.AutoMigrate(&models.Participation{}); err != nil {
		log.Fatal("Failed to migrate participations: ", err)
	}
	log.Println("Migrated participations!")

	log.Println("Migrating submissions...")
	if err := db.AutoMigrate(&models.Submission{}); err != nil {
		log.Fatal("Failed to migrate submissions: ", err)
	}
	log.Println("Migrated submissions!")

	log.Println("Migrating evaluations...")
	if err := db.AutoMigrate(&models.Evaluation{}); err != nil {
		log.Fatal("Failed to migrate evaluations: ", err)
	}
	log.Println("Migrated evaluations!")

	log.Println("Database migrated!")

	// Context for services
	ctx := context.Background()
	storageService := services.NewStorageService(ctx, os.Getenv("GCP_CREDS"))

	// Set up Gin router
	r := gin.New()
	r.Use(ginzero.Logger())

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Default route
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, gin-zerolog example")
		log.Println("Hello, gin-zerolog example")
	})

	// Swagger documentation
	docs.SwaggerInfo.Title = "Kiwi Collective API"
	docs.SwaggerInfo.Description = "API for the Kiwi Collective project."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	routes.UserRoutes(r, db, storageService)
	routes.SetupTeamRoutes(r, db)
	routes.HackathonRoutes(r, db)
	routes.FileRoutes(r, os.Getenv("GCP_CREDS"))
	routes.SubmissionRoutes(r, db, storageService)

	if err := seeders.SeedUsers(db); err != nil {
		log.Fatal("Failed to seed users: ", err)
	}
	if err := seeders.SeedHackathons(db); err != nil {
		log.Fatal("Failed to seed hackathons: ", err)
	}
	if err := seeders.SeedSkills(db); err != nil {
		log.Fatal("Failed to seed skills: ", err)
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
