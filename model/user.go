package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Email     string `json:"email" gorm:"primaryKey;uniqueIndex;not null" validate:"required,email"`
	Name      string `json:"name" gorm:"not null" validate:"required,min=4"`
	Password  string `json:"password" gorm:"not null" validate:"required,min=6"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
