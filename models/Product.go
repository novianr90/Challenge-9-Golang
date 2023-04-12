package models

import "time"

type Product struct {
	ID          uint   `gorm:"not null;primaryKey"`
	Name        string `gorm:"not null;type:varchar(100)"`
	Description string `gorm:"not null;type:varchar(100)"`
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
