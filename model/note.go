package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Id    string `gorm:"not null;uniqueIndex;primaryKey"`
	Title string `gorm:"not null"`
	Tags  string
	Note  string `gorm:"not null"`
	Email string `gorm:"not null"`
}
