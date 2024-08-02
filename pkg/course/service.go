package course

var courseService *Service

func InitCourseService(repository Repository) {
	courseService = NewCourseService(repository)
}

type Service struct {
	repo Repository
}

func NewCourseService(repository Repository) *Service {
	return &Service{repo: repository}
}

// pkg/course/service.go

func (s *Service) GetAllCourses(page, pageSize int) ([]Course, error) {
	return s.repo.GetAllCourses(page, pageSize)
}

func (s *Service) CreateCourse(course *Course) error {
	return s.repo.CreateCourse(course)
}

func (s *Service) UpdateCourse(course *Course) (*Course, error) {
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
	if course.Compact {
		existingCourse.Compact = course.Compact
	}
	if course.PitchAndPutt {
		existingCourse.PitchAndPutt = course.PitchAndPutt
	}
	if course.GolfID != "" {
		existingCourse.GolfID = course.GolfID
	}

	/*if err := s.repo.UpdateCourse(existingCourse); err != nil {
		return nil, err
	} */
	return existingCourse, nil
}

func (s *Service) DeleteCourse(id string) error {
	existingCourse, err := s.repo.GetCourseByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteCourseByID(existingCourse.ID)
}

func (s *Service) GetCourseByID(id string) (*Course, error) {
	return s.repo.GetCourseByID(id)
}

func (s *Service) GetCoursesByGolfID(golfID string, page, pageSize int) ([]Course, error) {
	return s.repo.GetCoursesByGolfID(golfID, page, pageSize)
}
