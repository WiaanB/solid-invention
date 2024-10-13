package model

import (
	"cinnanym/maps"
	"fmt"
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

func (u *User) ToMap(full bool) maps.Map {
	user := map[string]interface{}{
		"username": u.Username,
		"name":     u.Name,
		"email":    u.Email,
		"role":     u.Role,
	}

	if full {
		user["password"] = u.Password
	}

	return user
}

func (u *User) VerifyCreate() error {
	if u.Username == "" {
		return fmt.Errorf("missing field: username\n")
	}

	if u.Name == "" {
		return fmt.Errorf("missing field: name\n")
	}

	if u.Password == "" {
		return fmt.Errorf("missing field: password\n")
	}

	if u.Email == "" {
		return fmt.Errorf("missing field: email\n")
	}

	return nil
}
