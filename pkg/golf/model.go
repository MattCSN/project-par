package golf

import (
	"github.com/MattCSN/project-par/pkg/common"
	"github.com/MattCSN/project-par/pkg/utils"
	"github.com/jinzhu/gorm"
)

// Model represents a golf course
// @Description Model for a golf
type Model struct {
	common.Base
	Name           string  `json:"name" gorm:"type:varchar(255);not null" example:"Golf National"` // @Description Name of the golf course
	City           string  `json:"city" gorm:"type:varchar(255);not null" example:"Guyancourt"`    // @Description City where the golf course is located
	PostalCode     string  `json:"postalCode" gorm:"type:varchar(5);not null" example:"78280"`     // @Description Postal code of the golf course location
	Country        string  `json:"country" gorm:"type:varchar(255);not null" example:"France"`
	GoogleMapLinks string  `json:"googleMapLinks,omitempty" gorm:"type:varchar(255)" example:"https://maps.google.com/?q=Golf+National"` // @Description Google Maps link for the golf course
	Latitude       float64 `json:"latitude,omitempty" gorm:"type:double precision" example:"48.754"`                                     // @Description Latitude of the golf course
	Longitude      float64 `json:"longitude,omitempty" gorm:"type:double precision" example:"2.074"`                                     // @Description Longitude of the golf course
} // @name Golf

func (g *Model) TableName() string {
	return "golfs"
}

func (g *Model) BeforeSave(tx *gorm.DB) (err error) {
	var existingCount int64
	query := tx.Model(&Model{}).Where("name = ? AND city = ?", g.Name, g.City)
	if g.ID != "" {
		query = query.Not("id = ?", g.ID)
	}
	query.Count(&existingCount)

	if existingCount > 0 {
		return utils.ConflictError("Model 'name' and 'city' combination")
	}

	return nil
}
