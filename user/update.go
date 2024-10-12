package user

import (
	"cinnanym/database/surreal"
	"cinnanym/maps"
	"cinnanym/model"
	"errors"
	"fmt"
)

func Update(id string, user model.User) (maps.Map, error) {

	existingUser, err := surreal.FindOne(id)
	if err != nil {
		return nil, errors.New("that id does not exist")
	}

	fmt.Println("Updating", id, user)
	fmt.Println("Existing", existingUser)

	existingUser.Merge(user.ToMap())

	err = surreal.Update(surreal.UpdatePayload{ID: id, Data: existingUser})
	if err != nil {
		return nil, err
	}

	return existingUser, err
}
