package user

import (
	"cinnanym/database/surreal"
	"cinnanym/model"
	"errors"
)

func Create(username, name, password, email, role string) (model.User, error) {
	// ensure the required fields are there
	if username == "" || password == "" || email == "" {
		return model.User{}, errors.New("required fields are missing, try again")
	}

	// default role
	if role == "" {
		role = "user"
	}

	newUser := model.User{
		Username: username,
		Name:     name,
		Password: password,
		Email:    email,
		Role:     role,
	}

	err := surreal.Create(surreal.CreatePayload{
		Table:      "user",
		Identifier: "",
		Data:       newUser.ToMap(),
	})
	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}
