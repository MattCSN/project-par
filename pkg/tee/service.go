package tee

import "github.com/MattCSN/project-par/pkg/hole"

var teeService *Service

func InitTeeService(repository *Repository) {
	teeService = NewTeeService(repository)
}

type Service struct {
	repo *Repository
}

func NewTeeService(repository *Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) GetAllTees(page, pageSize int) ([]Model, error) {
	return s.repo.GetAllTees(page, pageSize)
}

func (s *Service) CreateTee(tee *Model) error {
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

func (s *Service) GetTeeByID(id string) (*Model, error) {
	return s.repo.GetTeeByID(id)
}

func (s *Service) GetTeesByHoleID(holeID string, page, pageSize int) ([]Model, error) {
	return s.repo.GetTeesByHoleID(holeID, page, pageSize)
}

func (s *Service) CreateTeesForCourse(courseID, color string) ([]Model, error) {
	holeService := hole.NewHoleService(hole.NewRepository())
	holes, err := holeService.GetHolesByCourseID(courseID, 1, 100)
	if err != nil {
		return nil, err
	}

	var createdTees []Model
	for _, holeModel := range holes {
		tee := Model{
			Color:  color,
			HoleID: holeModel.ID,
		}
		if err := s.repo.CreateTee(&tee); err != nil {
			return nil, err
		}
		createdTees = append(createdTees, tee)
	}

	return createdTees, nil
}

func (s *Service) UpdateTeesForCourse(courseID string, colors []string) ([]Model, error) {
	holeService := hole.NewHoleService(hole.NewRepository())
	holes, err := holeService.GetHolesByCourseID(courseID, 1, 100)
	if err != nil {
		return nil, err
	}

	var updatedTees []Model
	for _, holeModel := range holes {
		existingTees, err := s.repo.GetTeesByHoleID(holeModel.ID, 1, 100)
		if err != nil {
			return nil, err
		}

		existingColors := make(map[string]bool)
		for _, tee := range existingTees {
			existingColors[tee.Color] = true
		}

		for _, color := range colors {
			if !existingColors[color] {
				tee := Model{
					Color:  color,
					HoleID: holeModel.ID,
				}
				if err := s.repo.CreateTee(&tee); err != nil {
					return nil, err
				}
				updatedTees = append(updatedTees, tee)
			}
		}

		for _, tee := range existingTees {
			if !contains(colors, tee.Color) {
				if err := s.repo.DeleteTeeByID(tee.ID); err != nil {
					return nil, err
				}
			} else {
				updatedTees = append(updatedTees, tee)
			}
		}
	}

	return updatedTees, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
