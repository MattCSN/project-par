package common

import (
	"time"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}
