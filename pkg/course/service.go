package course

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/golf"
	"log"
)

var courseService *Service

func InitCourseService(repository *Repository, golfService *golf.Service) {
	courseService = NewCourseService(repository, golfService)
}

type Service struct {
	repo        *Repository
	golfService *golf.Service
}

func NewCourseService(repository *Repository, golfService *golf.Service) *Service {
	return &Service{repo: repository, golfService: golfService}
}

func (s *Service) GetAllCourses(page, pageSize int) ([]Model, error) {
	return s.repo.GetAllCourses(page, pageSize)
}

func (s *Service) CreateCourse(course *Model) error {
	log.Printf("Creating course with GolfID: %s", course.GolfID)
	if s.golfService == nil {
		log.Printf("Error: golfService is nil")
		return fmt.Errorf("golfService is nil")
	}
	if err := s.golfService.CheckGolfExists(course.GolfID); err != nil {
		log.Printf("Error checking if golf exists: %v", err)
		return err
	}
	log.Printf("Golf exists for GolfID: %s", course.GolfID)
	if err := s.repo.CreateCourse(course); err != nil {
		log.Printf("Error creating course: %v", err)
		return err
	}
	log.Printf("Course created successfully with ID: %s", course.ID)
	return nil
}

func (s *Service) UpdateCourse(course *Model) (*Model, error) {
	existingCourse, err := s.repo.GetCourseByID(course.ID)
	if err != nil {
		return nil, err
	}

	if course.Name != "" {
		existingCourse.Name = course.Name
	}
	if course.NumHoles != 0 {
		existingCourse.NumHoles = course.NumHoles
	}
	existingCourse.Compact = course.Compact
	existingCourse.PitchAndPutt = course.PitchAndPutt

	if err := s.repo.UpdateCourse(existingCourse); err != nil {
		return nil, err
	}
	return existingCourse, nil
}

func (s *Service) DeleteCourse(id string) error {
	existingCourse, err := s.repo.GetCourseByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteCourseByID(existingCourse.ID)
}

func (s *Service) GetCourseByID(id string) (*Model, error) {
	return s.repo.GetCourseByID(id)
}

func (s *Service) GetCoursesByGolfID(golfID string, page, pageSize int) ([]Model, error) {
	return s.repo.GetCoursesByGolfID(golfID, page, pageSize)
}

// Service method to get courses by multiple golf IDs
func (s *Service) GetCoursesByGolfIDs(golfIDs []string, page, pageSize int) ([]Model, error) {
	return s.repo.GetCoursesByGolfIDs(golfIDs, page, pageSize)
}
