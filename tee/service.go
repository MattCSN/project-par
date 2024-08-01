package tee

type Service struct {
	repo Repository
}

func NewTeeService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (gs *Service) GetAllTees() ([]Tee, error) {
	return gs.repo.GetAllTees()
}

func (gs *Service) CreateTee(tee *Tee) error {
	return gs.repo.CreateTee(tee)
}

func (gs *Service) UpdateTee(tee *Model) (*Model, error) {
	existingTee, err := gs.repo.GetTeeByID(tee.ID)
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

	err = gs.repo.UpdateTee(existingTee)
	if err != nil {
		return nil, err
	}
	return existingTee, nil
}

func (gs *Service) DeleteTee(id string) error {
	existingTee, err := gs.repo.GetTeeByID(id)
	if err != nil {
		return err
	}

	return gs.repo.DeleteTeeByID(existingTee.ID)
}

func (gs *Service) GetTeeByID(id string) (*Tee, error) {
	return gs.repo.GetTeeByID(id)
}
