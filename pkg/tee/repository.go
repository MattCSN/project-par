package tee

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/database"
	"github.com/MattCSN/project-par/pkg/utils"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) GetAllTees(page, pageSize int) ([]Model, error) {
	var tees []Model
	offset := (page - 1) * pageSize
	if err := database.DB.Offset(offset).Limit(pageSize).Find(&tees).Error; err != nil {
		return nil, err
	}
	return tees, nil
}

func (repo *Repository) CreateTee(tee *Model) error {
	return database.DB.Create(tee).Error
}

func (repo *Repository) GetTeeByID(id string) (*Model, error) {
	var tee Model
	err := database.DB.First(&tee, "id = ?", id).Error
	if err != nil {
		// TODO : Log the error here
		return nil, utils.NotFoundError(fmt.Sprintf("Model with id : %s", id))
	}
	return &tee, nil
}

func (repo *Repository) DeleteTeeByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Model{}).Error
}

func (repo *Repository) UpdateTee(tee *Model) error {
	return database.DB.Save(tee).Error
}

func (repo *Repository) GetTeesByHoleID(holeID string, page, pageSize int) ([]Model, error) {
	var tees []Model
	offset := (page - 1) * pageSize
	if err := database.DB.Where("hole_id = ?", holeID).Offset(offset).Limit(pageSize).Find(&tees).Error; err != nil {
		return nil, err
	}
	return tees, nil
}
