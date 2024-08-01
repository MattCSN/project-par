package main

import (
	"github.com/MattCSN/project-par/course"
	"github.com/MattCSN/project-par/database"
	_ "github.com/MattCSN/project-par/docs" // Import the generated docs
	"github.com/MattCSN/project-par/golf"
	"github.com/MattCSN/project-par/hole"
	"github.com/MattCSN/project-par/routes"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title ProjectPAR API
// @version 1.0
// @description This project is a comprehensive golf management system designed to facilitate the management of golf courses, including tracking of golf courses, holes, and tees.
// @host localhost:8080
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	database.Init(databaseURL, &golf.Golf{}, &course.Model{}, &hole.Model{}, &golf.Tee{})

	r := routes.SetupRouter()

	// Route pour la documentation Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
