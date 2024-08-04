package golf

import (
	"fmt"
	"github.com/MattCSN/project-par/pkg/database"
	"github.com/MattCSN/project-par/pkg/utils"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) GetAllGolfs(page, pageSize int) ([]Model, error) {
	var golfs []Model
	offset := (page - 1) * pageSize
	err := database.DB.Offset(offset).Limit(pageSize).Find(&golfs).Error
	return golfs, err
}

func (repo *Repository) CreateGolf(golf *Model) error {
	return database.DB.Create(golf).Error
}

func (repo *Repository) GetGolfByID(id string) (*Model, error) {
	var golf Model
	err := database.DB.First(&golf, "id = ?", id).Error
	if err != nil {
		return nil, utils.NotFoundError(fmt.Sprintf("Golf with id: %s", id))
	}
	return &golf, nil
}

func (repo *Repository) DeleteGolfByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Model{}).Error
}

func (repo *Repository) UpdateGolf(golf *Model) error {
	return database.DB.Save(golf).Error
}

func (repo *Repository) GolfExists(id string) (bool, error) {
	var count int64
	err := database.DB.Model(&Model{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (repo *Repository) SearchGolfs(query string) ([]Model, error) {
	var golfs []Model
	err := database.DB.Where("name ILIKE ?", "%"+query+"%").Find(&golfs).Error
	return golfs, err
}
