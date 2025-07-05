package database

import (
	"time"

	"github.com/google/uuid"
)

type ProductImages struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey;default:(UUID())"`
	DisplayOrder int       `gorm:"not null"`
	IsPrimary    bool      `gorm:"default:false"`
	ImageURL     string    `gorm:"type:text;not null"`
	ProductID    string    `gorm:"type:char(36);not null"`
	Product      Product   `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE;"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
