package course

import (
	"fmt"
	"github.com/MattCSN/project-par/database"
	"github.com/MattCSN/project-par/utils"
)

type Course = Model // Alias the Course type from the models package

type Repository interface {
	GetAllCourses() ([]Course, error)
	CreateCourse(course *Course) error
	GetCourseByID(id string) (*Course, error)
	DeleteCourseByID(id string) error
	UpdateCourse(course *Course) error
}

type courseRepository struct{}

func NewRepository() Repository {
	return &courseRepository{}
}

func (cr *courseRepository) GetAllCourses() ([]Course, error) {
	var golfs []Course
	return golfs, database.DB.Find(&golfs).Error
}

func (cr *courseRepository) CreateCourse(course *Course) error {
	return database.DB.Create(course).Error
}

func (cr *courseRepository) GetCourseByID(id string) (*Course, error) {
	var course Course
	err := database.DB.First(&course, "id = ?", id).Error
	if err != nil {
		// TODO : Log the error here
		return nil, utils.NotFoundError(fmt.Sprintf("Golf with id : %s", id))
	}
	return &course, nil
}

func (cr *courseRepository) DeleteCourseByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Course{}).Error
}

func (cr *courseRepository) UpdateCourse(course *Course) error {
	return database.DB.Save(course).Error
}
