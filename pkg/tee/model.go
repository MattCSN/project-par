package tee

import (
	"github.com/MattCSN/project-par/pkg/common"
)

// Model represents a tee on a golf tee
// @Description Model for a tee
type Model struct {
	common.Base
	Color    string `gorm:"type:varchar(50);not null"`
	Distance int    `gorm:"not null"`
	HoleID   string `gorm:"type:uuid;not null"`
} // @name Model

// TableName sets the insert table name for this struct type
func (Model) TableName() string {
	return "tees"
}
