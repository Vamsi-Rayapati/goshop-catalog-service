package database

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey;default:(UUID())"`
	Name        string    `gorm:"type:varchar(255);not null;"`
	Description string    `gorm:"type:text;"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Stock       int       `gorm:"type:int;not null;default:0"`
	CategoryID  uint      `gorm:"type:int;"`
	Category    Category  `gorm:"foreignKey:CategoryID;references:ID;constraint:OnDelete:SET NULL"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
