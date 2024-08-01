package hole

type Service struct {
	repo Repository
}

func NewHoleService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (gs *Service) GetAllHoles() ([]Hole, error) {
	return gs.repo.GetAllHoles()
}

func (gs *Service) CreateHole(hole *Hole) error {
	return gs.repo.CreateHole(hole)
}

func (gs *Service) UpdateHole(hole *Model) (*Model, error) {
	existingHole, err := gs.repo.GetHoleByID(hole.ID)
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

	err = gs.repo.UpdateHole(existingHole)
	if err != nil {
		return nil, err
	}
	return existingHole, nil
}

func (gs *Service) DeleteHole(id string) error {
	existingHole, err := gs.repo.GetHoleByID(id)
	if err != nil {
		return err
	}

	return gs.repo.DeleteHoleByID(existingHole.ID)
}

func (gs *Service) GetHoleByID(id string) (*Hole, error) {
	return gs.repo.GetHoleByID(id)
}
