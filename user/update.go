package user

import (
	"cinnanym/database/surreal"
	"cinnanym/model"
	"errors"
)

func Update(id string, user model.User) (model.User, error) {

	existingUser, err := surreal.FindOne[model.User](id)
	if err != nil {
		return model.User{}, errors.New("that id does not exist")
	}

	existingUser.UpdateValues(user)

	err = surreal.Update(surreal.UpdatePayload{ID: id, Data: existingUser.ToJson()})
	if err != nil {
		return model.User{}, err
	}

	return existingUser, err
}
