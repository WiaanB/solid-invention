package model

import (
	"github.com/surrealdb/surrealdb.go"
)

type User struct {
	surrealdb.Basemodel `table:"user"`
	Username            string `json:"username,omitempty"`
	Name                string `json:"name,omitempty"`
	Password            string `json:"password,omitempty"`
	Email               string `json:"email,omitempty"`
	Role                string `json:"role,omitempty"`
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
	if newValues.Email != "" {
		u.Email = newValues.Email
	}
	if newValues.Role != "" {
		u.Role = newValues.Role
	}
}

func (u *User) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"username": u.Username,
		"name":     u.Name,
		"email":    u.Email,
		"role":     u.Role,
	}
}
