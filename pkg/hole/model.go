package hole

import (
	"github.com/MattCSN/project-par/pkg/common"
)

// Model represents a hole on a golf course
// @Description Model for a hole
type Model struct {
	common.Base
	HoleNumber int    `gorm:"not null" example:"1"`                                              // @Description Hole number
	Par        int    `gorm:"not null" example:"4"`                                              // @Description Par for the hole
	CourseID   string `gorm:"type:uuid;not null" example:"123e4567-e89b-12d3-a456-426614174000"` // @Description ID of the course the hole is on
} // @name Hole

// TableName sets the insert table name for this struct type
func (Model) TableName() string {
	return "holes"
}
