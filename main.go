package main

import (
	"github.com/MattCSN/project-par/database"
	"github.com/MattCSN/project-par/golf"
	"github.com/MattCSN/project-par/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	database.Init(databaseURL, &golf.Golf{}, &golf.Course{}, &golf.Hole{}, &golf.Tee{})
	if err := routes.SetupRouter().Run(":8080"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
