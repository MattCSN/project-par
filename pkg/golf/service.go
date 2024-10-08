package golf

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/utils"
)

var golfService *Service

func InitGolfService(repository *Repository) {
	golfService = NewGolfService(repository)
}

type Service struct {
	repo *Repository
}

func NewGolfService(repository *Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) GetAllGolfs(page, pageSize int) ([]Model, error) {
	return s.repo.GetAllGolfs(page, pageSize)
}

func (s *Service) CreateGolf(golf *Model) error {
	return s.repo.CreateGolf(golf)
}

func (s *Service) UpdateGolf(golf *Model) (*Model, error) {
	existingGolf, err := s.repo.GetGolfByID(golf.ID)
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
	if golf.Country != "" {
		existingGolf.Country = golf.Country
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

	err = s.repo.UpdateGolf(existingGolf)
	if err != nil {
		return nil, err
	}
	return existingGolf, nil
}

func (s *Service) DeleteGolf(id string) error {
	existingGolf, err := s.repo.GetGolfByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteGolfByID(existingGolf.ID)
}

func (s *Service) GetGolfByID(id string) (*Model, error) {
	return s.repo.GetGolfByID(id)
}

func (s *Service) CheckGolfExists(golfID string) error {
	exists, err := s.repo.GolfExists(golfID)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NotFoundError(fmt.Sprintf("Golf with id: %s does not exist", golfID))
	}
	return nil
}

func (s *Service) SearchGolfs(searchTerm string) ([]Model, error) {
	return s.repo.SearchGolfs(searchTerm)
}

func (s *Service) SearchGolfsByProximity(longitude, latitude float64, page, pageSize int) ([]Model, error) {
	return s.repo.SearchGolfsByProximity(longitude, latitude, page, pageSize)
}
