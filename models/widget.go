package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Widget struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	PageID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	Type      string         `gorm:"not null"`
	Position  int
	Config    datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
