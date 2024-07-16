package database

import (
	"github.com/MattCSN/project-par/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import for its side effects
	"log"
)

var DB *gorm.DB

func Init(databaseURL string) {
	var err error
	DB, err = gorm.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	log.Println("Connected to the database successfully")

	if err = DB.AutoMigrate(&models.Golf{}, &models.Course{}, &models.Hole{}, &models.Tee{}).Error; err != nil {
		log.Fatalf("Error during migrations: %v", err)
	}
	log.Println("Database migrated successfully")
}
