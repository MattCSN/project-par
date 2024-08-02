package tee

import (
	"github.com/MattCSN/project-par/pkg/common"
)

// Model represents a tee on a golf tee
// @Description Model for a tee
type Model struct {
	common.Base
	Color    string `gorm:"type:varchar(50);not null" example:"White"`
	Distance int    `gorm:"not null" example:"353"`
	HoleID   string `gorm:"type:uuid;not null" example:"123e4567-e89b-12d3-a456-426614174000"`
} // @name Model

// TableName sets the insert table name for this struct type
func (Model) TableName() string {
	return "tees"
}
