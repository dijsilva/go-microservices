package entities

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Id          uint      `gorm:"primaryKey:Id"`
	ProfileName string    `json:"profile_name"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"json:"createdAt`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"json:"updatedAt`
	DeletedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"json:"deletedAt`
}
