package main

import (
	"challenges4/database"
	"challenges4/docs"
	"challenges4/models"
	"challenges4/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/madkins23/gin-utils/pkg/ginzero"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
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
	docs.SwaggerInfo.Title = "Kiwi Collective API"
	docs.SwaggerInfo.Description = "API for the Kiwi Collective project."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
	r := gin.New()
	r.Use(ginzero.Logger())

	//DB connection
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.ConnectDatabase()
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
