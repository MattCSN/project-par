package tee

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/database"
	"github.com/MattCSN/project-par/pkg/utils"
)

type Tee = Model

type Repository interface {
	GetAllTees(page, pageSize int) ([]Tee, error)
	CreateTee(*Tee) error
	GetTeeByID(id string) (*Tee, error)
	DeleteTeeByID(id string) error
	UpdateTee(tee *Tee) error
}

type teeRepository struct{}

func NewRepository() Repository {
	return &teeRepository{}
}

func (gr *teeRepository) GetAllTees(page, pageSize int) ([]Tee, error) {
	var tees []Tee
	offset := (page - 1) * pageSize
	if err := database.DB.Offset(offset).Limit(pageSize).Find(&tees).Error; err != nil {
		return nil, err
	}
	return tees, nil
}

func (gr *teeRepository) CreateTee(tee *Tee) error {
	return database.DB.Create(tee).Error
}

func (gr *teeRepository) GetTeeByID(id string) (*Tee, error) {
	var tee Tee
	err := database.DB.First(&tee, "id = ?", id).Error
	if err != nil {
		// TODO : Log the error here
		return nil, utils.NotFoundError(fmt.Sprintf("Tee with id : %s", id))
	}
	return &tee, nil
}

func (gr *teeRepository) DeleteTeeByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Tee{}).Error
}

func (gr *teeRepository) UpdateTee(tee *Tee) error {
	return database.DB.Save(tee).Error
}
