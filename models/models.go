package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
