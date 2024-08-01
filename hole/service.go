package hole

var holeService *Service

func InitHoleService(repository Repository) {
	holeService = NewHoleService(repository)
}

type Service struct {
	repo Repository
}

func NewHoleService(repository Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) GetAllHoles() ([]Hole, error) {
	return s.repo.GetAllHoles()
}

func (s *Service) CreateHole(hole *Hole) error {
	return s.repo.CreateHole(hole)
}

func (s *Service) UpdateHole(hole *Model) (*Model, error) {
	existingHole, err := s.repo.GetHoleByID(hole.ID)
	if err != nil {
		return nil, err
	}

	if hole.HoleNumber != 0 {
		existingHole.HoleNumber = hole.HoleNumber
	}
	if hole.Par != 0 {
		existingHole.Par = hole.Par
	}
	if hole.CourseID != "" {
		existingHole.CourseID = hole.CourseID
	}

	err = s.repo.UpdateHole(existingHole)
	if err != nil {
		return nil, err
	}
	return existingHole, nil
}

func (s *Service) DeleteHole(id string) error {
	existingHole, err := s.repo.GetHoleByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteHoleByID(existingHole.ID)
}

func (s *Service) GetHoleByID(id string) (*Hole, error) {
	return s.repo.GetHoleByID(id)
}
