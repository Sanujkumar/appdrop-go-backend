package models

import (
	"time"

	"github.com/google/uuid"
)
     
type Page struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null"`
	Route     string    `gorm:"unique;not null"`
	IsHome    bool      `gorm:"default:false"`
	BrandID   uuid.UUID `gorm:"type:uuid;index;not null"`
	Widgets   []Widget  `gorm:"foreignKey:PageID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


