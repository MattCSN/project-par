package hole

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/database"
	"github.com/MattCSN/project-par/pkg/utils"
)

type Hole = Model

type Repository interface {
	GetAllHoles(page, pageSize int) ([]Hole, error)
	CreateHole(*Hole) error
	GetHoleByID(id string) (*Hole, error)
	DeleteHoleByID(id string) error
	UpdateHole(hole *Hole) error
	GetHolesByCourseID(courseID string, page, pageSize int) ([]Hole, error)
}

type holeRepository struct{}

func NewRepository() Repository {
	return &holeRepository{}
}

func (gr *holeRepository) GetAllHoles(page, pageSize int) ([]Hole, error) {
	var holes []Hole
	offset := (page - 1) * pageSize
	if err := database.DB.Offset(offset).Limit(pageSize).Find(&holes).Error; err != nil {
		return nil, err
	}
	return holes, nil
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

func (gr *holeRepository) GetHolesByCourseID(courseID string, page, pageSize int) ([]Hole, error) {
	var holes []Hole
	offset := (page - 1) * pageSize
	if err := database.DB.Where("course_id = ?", courseID).Offset(offset).Limit(pageSize).Find(&holes).Error; err != nil {
		return nil, err
	}
	return holes, nil
}
