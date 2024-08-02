package hole

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/database"
	"github.com/MattCSN/project-par/pkg/utils"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) GetAllHoles(page, pageSize int) ([]Model, error) {
	var holes []Model
	offset := (page - 1) * pageSize
	if err := database.DB.Offset(offset).Limit(pageSize).Find(&holes).Error; err != nil {
		return nil, err
	}
	return holes, nil
}

func (repo *Repository) CreateHole(hole *Model) error {
	return database.DB.Create(hole).Error
}

func (repo *Repository) GetHoleByID(id string) (*Model, error) {
	var hole Model
	err := database.DB.First(&hole, "id = ?", id).Error
	if err != nil {
		// TODO : Log the error here
		return nil, utils.NotFoundError(fmt.Sprintf("Hole with id : %s", id))
	}
	return &hole, nil
}

func (repo *Repository) DeleteHoleByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Model{}).Error
}

func (repo *Repository) UpdateHole(hole *Model) error {
	return database.DB.Save(hole).Error
}

func (repo *Repository) GetHolesByCourseID(courseID string, page, pageSize int) ([]Model, error) {
	var holes []Model
	offset := (page - 1) * pageSize
	if err := database.DB.Where("course_id = ?", courseID).Offset(offset).Limit(pageSize).Find(&holes).Error; err != nil {
		return nil, err
	}
	return holes, nil
}
