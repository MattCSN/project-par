package course

import (
	"github.com/MattCSN/project-par/common"
)

// Model represents a golf course
// @Description Model for a golf course
type Model struct {
	common.Base
	Name         string `json:"name" gorm:"type:varchar(255);not null"`
	NumHoles     int    `json:"num_holes" gorm:"not null"`
	Compact      bool   `json:"compact" gorm:"default:false"`
	PitchAndPutt bool   `json:"pitch_and_putt" gorm:"default:false"`
	GolfID       string `json:"golf_id" gorm:"type:uuid;not null"`
} // @Name Course

// TableName sets the insert table name for this struct type
func (Model) TableName() string {
	return "courses"
}
