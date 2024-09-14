package user

import (
	"errors"
	"gorm.io/gorm"
	"gotcha/model"
)

func userEmailUnique(DB *gorm.DB, email string) bool {
	var existingUser model.User
	result := DB.Where("email = ?", email).First(&existingUser)
	if result.Error == nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return existingUser.Email != email
}
