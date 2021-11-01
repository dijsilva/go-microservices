package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:name`
	Email     string    `json: email`
	Username  string
	Password  string
	ProfileID int
	Profile   Profile   `gorm:"foreignKey:profile_id"`
	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp"json:"createdAt`
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp"json:"updatedAt`
	DeletedAt time.Time `gorm:"type:timestamp;default:current_timestamp"json:"deletedAt`
}
