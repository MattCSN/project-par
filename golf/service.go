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

func (gs *Service) AddGolfs(golfs []Golf) error {
	return gs.repo.AddGolfs(golfs)
}

func (gs *Service) UpdateGolf(golf *Golf) error {
	existingGolf, err := gs.repo.GetGolfByID(golf.ID)
	if err != nil {
		return err
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

	return gs.repo.UpdateGolf(existingGolf)
}

func (gs *Service) DeleteGolf(id string) error {
	return gs.repo.DeleteGolfByID(id)
}
