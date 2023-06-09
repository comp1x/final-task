package models

import (
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"time"
)

type Product struct {
	ID          uuid.UUID                    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string                       `gorm:"type:varchar(255);not null"`
	Description string                       `gorm:"type:varchar(255);not null"`
	Type        customer.CustomerProductType `gorm:"type:int;not null"`
	Weight      int32                        `gorm:"type:int;not null"`
	Price       float64                      `gorm:"type:float;not null"`
	CreatedAt   time.Time                    `gorm:"default:current_timestamp;not null"`
}
