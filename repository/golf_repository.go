package repository

import (
	"fmt"
	"github.com/MattCSN/project-par/database"
	"github.com/MattCSN/project-par/models"
	"github.com/MattCSN/project-par/utils"
)

type GolfRepository interface {
	GetAllGolfs() ([]models.Golf, error)
	CreateGolf(*models.Golf) error
	AddGolfs([]models.Golf) error
	GetGolfByID(id string) (*models.Golf, error)
	DeleteGolfByID(id string) error
	UpdateGolf(golf *models.Golf) error
}

type golfRepository struct{}

func NewGolfRepository() GolfRepository {
	return &golfRepository{}
}

func (gr *golfRepository) GetAllGolfs() ([]models.Golf, error) {
	var golfs []models.Golf
	return golfs, database.DB.Find(&golfs).Error
}

func (gr *golfRepository) CreateGolf(golf *models.Golf) error {
	return database.DB.Create(golf).Error
}

func (gr *golfRepository) AddGolfs(golfs []models.Golf) error {
	return database.DB.Create(&golfs).Error
}

func (gr *golfRepository) GetGolfByID(id string) (*models.Golf, error) {
	var golf models.Golf
	err := database.DB.First(&golf, "id = ?", id).Error
	if err != nil {
		// TODO : Log the error here
		return nil, utils.NotFoundError(fmt.Sprintf("Golf %s", id))
	}
	return &golf, nil
}

func (gr *golfRepository) DeleteGolfByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&models.Golf{}).Error
}

func (gr *golfRepository) UpdateGolf(golf *models.Golf) error {
	return database.DB.Save(golf).Error
}
