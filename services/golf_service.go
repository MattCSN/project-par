package services

import (
	"github.com/MattCSN/project-par/models"
	"github.com/MattCSN/project-par/repository"
)

type GolfService struct {
	repo repository.GolfRepository
}

func NewGolfService(repo repository.GolfRepository) *GolfService {
	return &GolfService{repo}
}

func (gs *GolfService) GetAllGolfs() ([]models.Golf, error) {
	return gs.repo.GetAllGolfs()
}

func (gs *GolfService) CreateGolf(golf *models.Golf) error {
	return gs.repo.CreateGolf(golf)
}

func (gs *GolfService) AddGolfs(golfs []models.Golf) error {
	return gs.repo.AddGolfs(golfs)
}

func (gs *GolfService) UpdateGolf(golf *models.Golf) error {
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

func (gs *GolfService) DeleteGolf(id string) error {
	return gs.repo.DeleteGolfByID(id)
}
