package golf

import (
	"github.com/MattCSN/project-par/common"
	"github.com/MattCSN/project-par/utils"
	"github.com/jinzhu/gorm"
)

type Golf struct {
	common.Base
	Name           string  `json:"name" gorm:"type:varchar(255);not null"`
	City           string  `json:"city" gorm:"type:varchar(255);not null"`
	PostalCode     string  `json:"postalCode" gorm:"type:varchar(5);not null"`
	GoogleMapLinks string  `json:"googleMapLinks,omitempty" gorm:"type:varchar(255)"`
	Latitude       float64 `json:"latitude,omitempty" gorm:"type:double precision"`
	Longitude      float64 `json:"longitude,omitempty" gorm:"type:double precision"`
}

type Hole struct {
	common.Base
	HoleNumber int    `gorm:"not null"`
	Par        int    `gorm:"not null"`
	CourseID   string `gorm:"type:uuid;not null"`
}

type Tee struct {
	common.Base
	Color    string `gorm:"type:varchar(50);not null"`
	Distance int    `gorm:"not null"`
	HoleID   string `gorm:"type:uuid;not null"`
}

func (g *Golf) TableName() string {
	return "golfs"
}

func (g *Golf) BeforeSave(tx *gorm.DB) (err error) {
	var existingCount int64
	query := tx.Model(&Golf{}).Where("name = ? AND city = ?", g.Name, g.City)
	if g.ID != "" {
		query = query.Not("id = ?", g.ID)
	}
	query.Count(&existingCount)

	if existingCount > 0 {
		return utils.ConflictError("Golf 'name' and 'city' combination")
	}

	return nil
}
