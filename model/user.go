package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique,not null"`
	Name     string
	Password string `gorm:"unique,not null"`
	Email    string `gorm:"unique,not null"`
	Role     string `gorm:"default:'user'"`
}
