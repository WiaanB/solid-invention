package model

import (
	"cinnanym/maps"
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

func (u *User) ToMap() maps.Map {
	return map[string]interface{}{
		"username": u.Username,
		"name":     u.Name,
		"email":    u.Email,
		"role":     u.Role,
	}
}
