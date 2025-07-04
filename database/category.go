package database

import "time"

type Category struct {
	ID        uint      `gorm:"type:int;primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
