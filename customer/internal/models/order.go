package models

import (
	"github.com/google/uuid"
)

type Order struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
}