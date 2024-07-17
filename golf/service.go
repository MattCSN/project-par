package golf

type Service struct {
	repo Repository
}

func NewGolfService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (gs *Service) GetAllGolfs() ([]Golf, error) {
	return gs.repo.GetAllGolfs()
}

func (gs *Service) CreateGolf(golf *Golf) error {
	return gs.repo.CreateGolf(golf)
}

func (gs *Service) UpdateGolf(golf *Golf) (*Golf, error) {
	existingGolf, err := gs.repo.GetGolfByID(golf.ID)
	if err != nil {
		return nil, err
	}

	if golf.Name != "" {
		existingGolf.Name = golf.Name
	}
	if golf.City != "" {
		existingGolf.City = golf.City
	}
	if golf.PostalCode != "" {
		existingGolf.PostalCode = golf.PostalCode
	}
	if golf.GoogleMapLinks != "" {
		existingGolf.GoogleMapLinks = golf.GoogleMapLinks
	}
	if golf.Latitude != 0 {
		existingGolf.Latitude = golf.Latitude
	}
	if golf.Longitude != 0 {
		existingGolf.Longitude = golf.Longitude
	}

	err = gs.repo.UpdateGolf(existingGolf)
	if err != nil {
		return nil, err
	}
	return existingGolf, nil
}

func (gs *Service) DeleteGolf(id string) error {
	existingGolf, err := gs.repo.GetGolfByID(id)
	if err != nil {
		return err
	}

	return gs.repo.DeleteGolfByID(existingGolf.ID)
}

func (gs *Service) GetGolfByID(id string) (*Golf, error) {
	return gs.repo.GetGolfByID(id)
}
