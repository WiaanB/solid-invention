package user

import (
	"errors"
	"gorm.io/gorm"
	"gotcha/model"
)

func CreateNewUser(DB *gorm.DB, username, name, password, email, role string) (error, model.User) {
	// ensure the required fields are there
	if username == "" || password == "" || email == "" {
		return errors.New("required fields are missing. big mistake, buddy"), model.User{}
	}

	// ensure the email is unique
	if !userEmailUnique(DB, email) {
		return errors.New("email already belongs to one of these idiots"), model.User{}
	}

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

func SaveUser(DB *gorm.DB, user model.User) error {
	// update the user if it is found, otherwise create it
	var existingUser model.User
	result := DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
	}

	if existingUser.Email == user.Email {
		existingUser.UpdateValues(user)
		result = DB.Save(&existingUser)
		return result.Error
	}

	result = DB.Create(&user)
	return result.Error
}
