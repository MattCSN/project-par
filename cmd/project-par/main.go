package main

import (
	"github.com/MattCSN/project-par/pkg/jsonimport"
	"github.com/MattCSN/project-par/pkg/view"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/MattCSN/project-par/docs"
	"github.com/MattCSN/project-par/pkg/course"
	"github.com/MattCSN/project-par/pkg/database"
	"github.com/MattCSN/project-par/pkg/golf"
	"github.com/MattCSN/project-par/pkg/hole"
	"github.com/MattCSN/project-par/pkg/router"
	"github.com/MattCSN/project-par/pkg/tee"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Project-PAR API
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
	course.InitCourseService(course.NewRepository(), &golf.Service{})
	hole.InitHoleService(hole.NewRepository())
	tee.InitTeeService(tee.NewRepository())
	view.InitViewService()

	activeJSONImport, err := strconv.ParseBool(os.Getenv("ACTIVE_JSON_IMPORT"))
	if err != nil {
		log.Fatalf("Error parsing ACTIVE_JSON_IMPORT: %v", err)
	}
	if activeJSONImport {
		jsonimport.ImportGolfData("pkg/jsonimport/golf2import.json")
	}

	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}

// Exported function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	main()
}
