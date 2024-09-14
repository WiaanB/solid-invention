package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `gorm:"not null" json:"password,omitempty"`
	Email    string `gorm:"unique,not null" json:"email,omitempty"`
	Role     string `gorm:"default:'user'" json:"role,omitempty"`
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *User) UpdateValues(newValues User) {
	if newValues.Username != "" {
		u.Username = newValues.Username
	}
	if newValues.Name != "" {
		u.Name = newValues.Name
	}
	if newValues.Password != "" {
		u.Password = newValues.Password
	}
	if newValues.Email != "" {
		u.Email = newValues.Email
	}
	if newValues.Role != "" {
		u.Role = newValues.Role
	}
}
