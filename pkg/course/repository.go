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

func (repo *Repository) GetCoursesByGolfIDs(golfIDs []string, page, pageSize int) ([]Model, error) {
	var courses []Model
	offset := (page - 1) * pageSize
	if err := database.DB.Where("golf_id IN (?)", golfIDs).Offset(offset).Limit(pageSize).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (repo *Repository) GetCourseDetails(offset, limit int) ([]DetailsDTO, error) {
	var courseDetails []DetailsDTO
	query := `
		SELECT c.id, c.created_at, c.updated_at, c.name as course_name, g.name as golf_name, c.num_holes, c.pitch_and_putt, c.compact, g.postal_code, g.city, g.country
		FROM courses c
		JOIN golfs g ON c.golf_id = g.id
		LIMIT ? OFFSET ?
	`
	if err := database.DB.Raw(query, limit, offset).Scan(&courseDetails).Error; err != nil {
		return nil, err
	}
	return courseDetails, nil
}

func (repo *Repository) SearchCourseDetails(searchTerm string, page, pageSize int) ([]DetailsDTO, error) {
	var courseDetails []DetailsDTO
	offset := (page - 1) * pageSize
	query := `
		SELECT c.id, c.created_at, c.updated_at, c.name as course_name, g.name as golf_name, c.num_holes, c.pitch_and_putt, c.compact, g.postal_code, g.city, g.country
		FROM courses c
		JOIN golfs g ON c.golf_id = g.id
		WHERE c.name ILIKE ? OR g.name ILIKE ? OR g.city ILIKE ? OR g.postal_code ILIKE ?
		LIMIT ? OFFSET ?
	`
	if err := database.DB.Raw(query, "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%", pageSize, offset).Scan(&courseDetails).Error; err != nil {
		return nil, err
	}
	if len(courseDetails) == 0 {
		return nil, utils.NotFoundError("No courses found matching the search criteria")
	}
	return courseDetails, nil
}

func (repo *Repository) SearchCourseDetailsByProximity(longitude, latitude float64, page, pageSize int) ([]DetailsDTO, error) {
	var courseDetails []DetailsDTO
	offset := (page - 1) * pageSize
	query := `
		SELECT c.id, c.created_at, c.updated_at, c.name as course_name, g.name as golf_name, c.num_holes, c.pitch_and_putt, c.compact, g.postal_code, g.city, g.country,
		( 6371 * acos( cos( radians(?) ) * cos( radians( g.latitude ) ) * cos( radians( g.longitude ) - radians(?) ) + sin( radians(?) ) * sin( radians( g.latitude ) ) ) ) AS distance 
		FROM courses c
		JOIN golfs g ON c.golf_id = g.id
		ORDER BY distance
		LIMIT ? OFFSET ?
	`
	if err := database.DB.Raw(query, latitude, longitude, latitude, pageSize, offset).Scan(&courseDetails).Error; err != nil {
		return nil, err
	}
	if len(courseDetails) == 0 {
		return nil, utils.NotFoundError("No courses found near the provided coordinates")
	}
	return courseDetails, nil
}
