package user

import (
	"errors"
	"gotcha/model"
)

func CreateNewUser(username, name, password, email, role string) (error, model.User) {
	// ensure the required fields are there
	if username == "" || password == "" || email == "" {
		return errors.New("required fields are missing. big mistake, buddy"), model.User{}
	}

	// TODO: ensure the email is unique

	// default role
	if role == "" {
		role = "USER"
	}

	return nil, model.User{
		Username: username,
		Name:     name,
		Password: password,
		Email:    email,
		Role:     role,
	}
}

// TODO: consider if saving should be user or DB based, leaning towards saving on the model instead
func UpdateUser(user model.User, newValues model.User) error {
	// ensure the required fields are there
	if newValues.Email == "" {
		return errors.New("you need an email to update the user by")
	}

	// TODO: ensure the email is unique

	user.UpdateValues(newValues)
	// TODO: handle the new values

	return nil
}
