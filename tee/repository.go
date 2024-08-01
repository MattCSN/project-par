package tee

import (
	"fmt"
	"github.com/MattCSN/project-par/database"
	"github.com/MattCSN/project-par/utils"
)

type Tee = Model

type Repository interface {
	GetAllTees() ([]Tee, error)
	CreateTee(*Tee) error
	GetTeeByID(id string) (*Tee, error)
	DeleteTeeByID(id string) error
	UpdateTee(tee *Tee) error
}

type teeRepository struct{}

func NewTeeRepository() Repository {
	return &teeRepository{}
}

func (gr *teeRepository) GetAllTees() ([]Tee, error) {
	var tees []Tee
	return tees, database.DB.Find(&tees).Error
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
