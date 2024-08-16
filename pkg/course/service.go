package course

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/golf"
	"github.com/MattCSN/project-par/pkg/hole"
	"github.com/MattCSN/project-par/pkg/tee"
	"log"
)

var courseService *Service

func InitCourseService(repository *Repository, golfService *golf.Service, holeService *hole.Service, teeService *tee.Service) {
	courseService = NewCourseService(repository, golfService, holeService, teeService)
}

type Service struct {
	repo        *Repository
	golfService *golf.Service
	holeService *hole.Service
	teeService  *tee.Service
}

func NewCourseService(repository *Repository, golfService *golf.Service, holeService *hole.Service, teeService *tee.Service) *Service {
	return &Service{repo: repository, golfService: golfService, holeService: holeService, teeService: teeService}
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

func (s *Service) GetCoursesByGolfIDs(golfIDs []string, page, pageSize int) ([]Model, error) {
	return s.repo.GetCoursesByGolfIDs(golfIDs, page, pageSize)
}

func (s *Service) GetCourseDetails(page, pageSize int) ([]DetailsDTO, error) {
	offset := (page - 1) * pageSize
	return s.repo.GetCourseDetails(offset, pageSize)
}

func (s *Service) SearchCourseDetails(searchTerm string, page, pageSize int) ([]DetailsDTO, error) {
	return s.repo.SearchCourseDetails(searchTerm, page, pageSize)
}

func (s *Service) SearchCourseDetailsByProximity(longitude, latitude float64, page, pageSize int) ([]DetailsDTO, error) {
	return s.repo.SearchCourseDetailsByProximity(longitude, latitude, page, pageSize)
}

func (s *Service) GetGolfAndCourseDetails(courseID string) (CourseDetailsDTO, error) {
	// Fetch the course details
	course, err := s.repo.GetCourseByID(courseID)
	if err != nil {
		return CourseDetailsDTO{}, err
	}

	// Fetch the golfDetails details
	golfDetails, err := s.golfService.GetGolfByID(course.GolfID)
	if err != nil {
		return CourseDetailsDTO{}, err
	}

	// Fetch all courses for the golfDetails
	courses, err := s.repo.GetCoursesByGolfID(course.GolfID, 1, 100)
	if err != nil {
		return CourseDetailsDTO{}, err
	}

	// Fetch holes for the specific course
	holes, err := s.holeService.GetHolesByCourseID(courseID, 1, 100)
	if err != nil {
		return CourseDetailsDTO{}, err
	}

	// Fetch tees for each holeDetail
	var holeDetails []HoleDetailsDTO
	for _, holeDetail := range holes {
		tees, err := s.teeService.GetTeesByHoleID(holeDetail.ID, 1, 100)
		if err != nil {
			return CourseDetailsDTO{}, err
		}
		holeDetails = append(holeDetails, newHoleDetailsDTO(holeDetail, tees))
	}

	// Prepare the response DTO using newCourseDetailsDTO
	courseDetails := newCourseDetailsDTO(*course, courses, *golfDetails, holeDetails)

	return courseDetails, nil
}
