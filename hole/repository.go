package hole

import (
	"fmt"
	"github.com/MattCSN/project-par/database"
	"github.com/MattCSN/project-par/utils"
)

type Hole = Model

type Repository interface {
	GetAllHoles() ([]Hole, error)
	CreateHole(*Hole) error
	GetHoleByID(id string) (*Hole, error)
	DeleteHoleByID(id string) error
	UpdateHole(hole *Hole) error
}

type holeRepository struct{}

func NewRepository() Repository {
	return &holeRepository{}
}

func (gr *holeRepository) GetAllHoles() ([]Hole, error) {
	var holes []Hole
	return holes, database.DB.Find(&holes).Error
}

func (gr *holeRepository) CreateHole(hole *Hole) error {
	return database.DB.Create(hole).Error
}

func (gr *holeRepository) GetHoleByID(id string) (*Hole, error) {
	var hole Hole
	err := database.DB.First(&hole, "id = ?", id).Error
	if err != nil {
		// TODO : Log the error here
		return nil, utils.NotFoundError(fmt.Sprintf("Hole with id : %s", id))
	}
	return &hole, nil
}

func (gr *holeRepository) DeleteHoleByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Hole{}).Error
}

func (gr *holeRepository) UpdateHole(hole *Hole) error {
	return database.DB.Save(hole).Error
}
