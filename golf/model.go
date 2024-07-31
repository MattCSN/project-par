package golf

import (
	"github.com/MattCSN/project-par/common"
	"github.com/MattCSN/project-par/utils"
	"github.com/jinzhu/gorm"
)

// Golf represents a golf course
// @Description Model for a golf
type Golf struct {
	common.Base
	Name           string  `json:"name" gorm:"type:varchar(255);not null" example:"Pebble Beach"`                                       // @Description Name of the golf course
	City           string  `json:"city" gorm:"type:varchar(255);not null" example:"Monterey"`                                           // @Description City where the golf course is located
	PostalCode     string  `json:"postalCode" gorm:"type:varchar(5);not null" example:"93953"`                                          // @Description Postal code of the golf course location
	GoogleMapLinks string  `json:"googleMapLinks,omitempty" gorm:"type:varchar(255)" example:"https://maps.google.com/?q=Pebble+Beach"` // @Description Google Maps link for the golf course
	Latitude       float64 `json:"latitude,omitempty" gorm:"type:double precision" example:"36.567"`                                    // @Description Latitude of the golf course
	Longitude      float64 `json:"longitude,omitempty" gorm:"type:double precision" example:"-121.950"`                                 // @Description Longitude of the golf course
} // @name Golf

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
