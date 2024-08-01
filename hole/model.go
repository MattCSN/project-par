package hole

import (
	"github.com/MattCSN/project-par/common"
)

// Model represents a hole on a golf course
// @Description Model for a hole
type Model struct {
	common.Base
	HoleNumber int    `gorm:"not null"`
	Par        int    `gorm:"not null"`
	CourseID   string `gorm:"type:uuid;not null"`
} // @name Hole
