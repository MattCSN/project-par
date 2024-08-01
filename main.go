package main

import (
	"log"
	"os"

	"github.com/MattCSN/project-par/course"
	"github.com/MattCSN/project-par/database"
	_ "github.com/MattCSN/project-par/docs"
	"github.com/MattCSN/project-par/golf"
	"github.com/MattCSN/project-par/hole"
	"github.com/MattCSN/project-par/router"
	"github.com/MattCSN/project-par/tee"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title ProjectPAR API
// @version 1.0.0-Beta
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

	database.Init(databaseURL, &golf.Model{}, &course.Model{}, &hole.Model{}, &tee.Model{})

	golf.InitGolfService(golf.NewRepository())
	course.InitCourseService(course.NewRepository())
	hole.InitHoleService(hole.NewRepository())
	tee.InitTeeService(tee.NewRepository())

	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
