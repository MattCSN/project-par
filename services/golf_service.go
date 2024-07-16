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
