package repository

import (
	"github.com/MattCSN/project-par/database"
	"github.com/MattCSN/project-par/models"
)

type GolfRepository interface {
	GetAllGolfs() ([]models.Golf, error)
	CreateGolf(*models.Golf) error
	AddGolfs([]models.Golf) error
}

type golfRepository struct{}

func NewGolfRepository() GolfRepository {
	return &golfRepository{}
}

func (gr *golfRepository) GetAllGolfs() ([]models.Golf, error) {
	var golfs []models.Golf
	err := database.DB.Find(&golfs).Error
	return golfs, err
}

func (gr *golfRepository) CreateGolf(golf *models.Golf) error {
	return database.DB.Create(golf).Error
}

func (gr *golfRepository) AddGolfs(golfs []models.Golf) error {
	for _, golf := range golfs {
		if err := database.DB.Create(&golf).Error; err != nil {
			return err
		}
	}
	return nil
}
