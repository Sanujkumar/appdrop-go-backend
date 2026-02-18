package models

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"unique;not null"`
	Pages     []Page    `gorm:"foreignKey:BrandID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}