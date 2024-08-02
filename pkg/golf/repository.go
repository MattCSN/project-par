package golf

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/database"
	"github.com/MattCSN/project-par/pkg/utils"
)

type Golf = Model

type Repository interface {
	GetAllGolfs(page, pageSize int) ([]Golf, error)
	CreateGolf(*Golf) error
	GetGolfByID(id string) (*Golf, error)
	DeleteGolfByID(id string) error
	UpdateGolf(golf *Golf) error
}

type golfRepository struct{}

func NewRepository() Repository {
	return &golfRepository{}
}

func (repo *golfRepository) GetAllGolfs(page, pageSize int) ([]Golf, error) {
	var golfs []Golf
	offset := (page - 1) * pageSize
	if err := database.DB.Offset(offset).Limit(pageSize).Find(&golfs).Error; err != nil {
		return nil, err
	}
	return golfs, nil
}

func (repo *golfRepository) CreateGolf(golf *Golf) error {
	return database.DB.Create(golf).Error
}

func (repo *golfRepository) GetGolfByID(id string) (*Golf, error) {
	var golf Golf
	err := database.DB.First(&golf, "id = ?", id).Error
	if err != nil {
		return nil, utils.NotFoundError(fmt.Sprintf("Golf with id: %s", id))
	}
	return &golf, nil
}

func (repo *golfRepository) DeleteGolfByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Golf{}).Error
}

func (repo *golfRepository) UpdateGolf(golf *Golf) error {
	return database.DB.Save(golf).Error
}
