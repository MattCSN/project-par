package golf

import (
	"github.com/MattCSN/project-par/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type Golf struct {
	ID             string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name           string    `json:"name" gorm:"type:varchar(255);not null"`
	City           string    `json:"city" gorm:"type:varchar(255);not null"`
	PostalCode     string    `json:"postalCode" gorm:"type:varchar(5);not null"`
	GoogleMapLinks string    `json:"googleMapLinks,omitempty" gorm:"type:varchar(255)"`
	Latitude       float64   `json:"latitude" gorm:"type:double precision"`
	Longitude      float64   `json:"longitude" gorm:"type:double precision"`
	CreatedAt      time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}

func (g *Golf) TableName() string {
	return "golfs"
}

func (g *Golf) BeforeSave(tx *gorm.DB) (err error) {
	// Check for an existing entity with the same Name and City
	var existingCount int64
	query := tx.Model(&Golf{}).Where("name = ? AND city = ?", g.Name, g.City)
	if g.ID != "" {
		query = query.Not("id = ?", g.ID) // Exclude the current entity if updating
	}
	query.Count(&existingCount)

	if existingCount > 0 {
		return utils.ConflictError("Golf 'name' and 'city' combination")
	}

	return nil
}

type Course struct {
	ID           string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string    `gorm:"type:varchar(255);not null"`
	NumHoles     int       `gorm:"not null"`
	Compact      bool      `gorm:"default:false"`
	PitchAndPutt bool      `gorm:"default:false"`
	GolfID       string    `gorm:"type:uuid;not null"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Hole struct {
	ID         string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	HoleNumber int       `gorm:"not null"`
	Par        int       `gorm:"not null"`
	CourseID   string    `gorm:"type:uuid;not null"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Tee struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Color     string    `gorm:"type:varchar(50);not null"`
	Distance  int       `gorm:"not null"`
	HoleID    string    `gorm:"type:uuid;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
