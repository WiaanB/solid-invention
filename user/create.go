package user

import (
	"cinnanym/model"
	"errors"
	"github.com/surrealdb/surrealdb.go"
)

func Create(DB *surrealdb.DB, username, name, password, email, role string) (model.User, error) {
	// ensure the required fields are there
	if username == "" || password == "" || email == "" {
		return model.User{}, errors.New("required fields are missing, try again")
	}

	_, err := FindUser(DB, email)
	if err == nil {
		return model.User{}, errors.New("user already exists")
	}

	// default role
	if role == "" {
		role = "USER"
	}

	newUser := model.User{
		Username: username,
		Name:     name,
		Password: password,
		Email:    email,
		Role:     role,
	}

	err = CreateUser(DB, newUser)
	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}
