package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"not null;primaryKey"`
	Email     string `gorm:"not null;type:varchar(100)"`
	Password  string `gorm:"not null;type:varchar(100)"`
	Username  string `gorm:"type:varchar(100)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var user User
	err := tx.Where("email", u.Email).First(&user).Error

	if err == nil {
		return errors.New("already registered")
	}

	return nil
}
