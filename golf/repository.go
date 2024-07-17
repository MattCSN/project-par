package golf

import (
	"fmt"
	"github.com/MattCSN/project-par/database"
	"github.com/MattCSN/project-par/utils"
)

type Repository interface {
	GetAllGolfs() ([]Golf, error)
	CreateGolf(*Golf) error
	AddGolfs([]Golf) error
	GetGolfByID(id string) (*Golf, error)
	DeleteGolfByID(id string) error
	UpdateGolf(golf *Golf) error
}

type golfRepository struct{}

func NewGolfRepository() Repository {
	return &golfRepository{}
}

func (gr *golfRepository) GetAllGolfs() ([]Golf, error) {
	var golfs []Golf
	return golfs, database.DB.Find(&golfs).Error
}

func (gr *golfRepository) CreateGolf(golf *Golf) error {
	return database.DB.Create(golf).Error
}

func (gr *golfRepository) AddGolfs(golfs []Golf) error {
	return database.DB.Create(&golfs).Error
}

func (gr *golfRepository) GetGolfByID(id string) (*Golf, error) {
	var golf Golf
	err := database.DB.First(&golf, "id = ?", id).Error
	if err != nil {
		// TODO : Log the error here
		return nil, utils.NotFoundError(fmt.Sprintf("Golf with id : %s", id))
	}
	return &golf, nil
}

func (gr *golfRepository) DeleteGolfByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Golf{}).Error
}

func (gr *golfRepository) UpdateGolf(golf *Golf) error {
	return database.DB.Save(golf).Error
}
