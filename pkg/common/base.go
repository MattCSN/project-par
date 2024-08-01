package common

import (
	"time"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key" example:"123e4567-e89b-12d3-a456-426614174000 (auto-generated)"` // @Description ID of the model, automatically generated
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP" example:"2021-01-01T00:00:00Z (auto-generated)"`                                 // @Description Creation date of the model, automatically generated
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP" example:"2021-01-01T00:00:00Z (auto-generated)"`                                 // @Description Last update date of the model, automatically generated
}
