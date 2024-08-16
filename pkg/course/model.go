package course

import (
	"github.com/MattCSN/project-par/pkg/common"
)

// Model represents a golf course
// @Description Model for a golf course
type Model struct {
	common.Base
	Name         string `json:"name" gorm:"type:varchar(255);not null" example:"L’Albatros"`
	NumHoles     int    `json:"numberHoles" gorm:"not null" example:"18"`
	Compact      bool   `json:"compact" gorm:"default:false" example:"false"`
	PitchAndPutt bool   `json:"pitchAndPutt" gorm:"default:false" example:"false"`
	GolfID       string `json:"golfID" gorm:"type:uuid;not null" example:"123e4567-e89b-12d3-a456-426614174000"`
} // @Name Course

// TableName sets the insert table name for this struct type
func (Model) TableName() string {
	return "courses"
}
