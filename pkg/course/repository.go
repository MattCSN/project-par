package course

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/database"
	"github.com/MattCSN/project-par/pkg/utils"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) GetCoursesByGolfID(golfID string, page, pageSize int) ([]Model, error) {
	var courses []Model
	offset := (page - 1) * pageSize
	if err := database.DB.Where("golf_id = ?", golfID).Offset(offset).Limit(pageSize).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (repo *Repository) GetAllCourses(page, pageSize int) ([]Model, error) {
	var courses []Model
	offset := (page - 1) * pageSize
	if err := database.DB.Offset(offset).Limit(pageSize).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (repo *Repository) CreateCourse(course *Model) error {
	return database.DB.Create(course).Error
}

func (repo *Repository) GetCourseByID(id string) (*Model, error) {
	var course Model
	err := database.DB.First(&course, "id = ?", id).Error
	if err != nil {
		return nil, utils.NotFoundError(fmt.Sprintf("Course with id : %s", id))
	}
	return &course, nil
}

func (repo *Repository) DeleteCourseByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Model{}).Error
}

func (repo *Repository) UpdateCourse(course *Model) error {
	return database.DB.Save(course).Error
}
