package course

type Service struct {
	repo Repository
}

func NewCourseService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (cs *Service) GetAllCourses() ([]Course, error) {
	return cs.repo.GetAllCourses()
}

func (cs *Service) CreateCourse(course *Course) error {
	return cs.repo.CreateCourse(course)
}

func (cs *Service) UpdateCourse(course *Course) (*Course, error) {
	existingCourse, err := cs.repo.GetCourseByID(course.ID)
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

	/*if err := cs.repo.UpdateCourse(existingCourse); err != nil {
		return nil, err
	} */
	return existingCourse, nil
}

func (cs *Service) DeleteCourse(id string) error {
	existingCourse, err := cs.repo.GetCourseByID(id)
	if err != nil {
		return err
	}

	return cs.repo.DeleteCourseByID(existingCourse.ID)
}

func (cs *Service) GetCourseByID(id string) (*Course, error) {
	return cs.repo.GetCourseByID(id)
}
