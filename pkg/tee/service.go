package tee

var teeService *Service

func InitTeeService(repository Repository) {
	teeService = NewTeeService(repository)
}

type Service struct {
	repo Repository
}

func NewTeeService(repository Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) GetAllTees(page, pageSize int) ([]Tee, error) {
	return s.repo.GetAllTees(page, pageSize)
}

func (s *Service) CreateTee(tee *Tee) error {
	return s.repo.CreateTee(tee)
}

func (s *Service) UpdateTee(tee *Model) (*Model, error) {
	existingTee, err := s.repo.GetTeeByID(tee.ID)
	if err != nil {
		return nil, err
	}

	if tee.Color != "" {
		existingTee.Color = tee.Color
	}
	if tee.Distance != 0 {
		existingTee.Distance = tee.Distance
	}
	if tee.HoleID != "" {
		existingTee.HoleID = tee.HoleID
	}

	err = s.repo.UpdateTee(existingTee)
	if err != nil {
		return nil, err
	}
	return existingTee, nil
}

func (s *Service) DeleteTee(id string) error {
	existingTee, err := s.repo.GetTeeByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteTeeByID(existingTee.ID)
}

func (s *Service) GetTeeByID(id string) (*Tee, error) {
	return s.repo.GetTeeByID(id)
}
