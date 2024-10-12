package restaurant

import (
	"cinnanym/database/surreal"
	"cinnanym/maps"
	"cinnanym/model"
	"errors"
)

func Update(id string, restaurant model.Restaurant) (maps.Map, error) {

	existingRestaurant, err := surreal.FindOne(id)
	if err != nil {
		return nil, errors.New("that id does not exist")
	}

	existingRestaurant.Merge(restaurant.ToMap())

	err = surreal.Update(surreal.UpdatePayload{ID: id, Data: existingRestaurant})
	if err != nil {
		return nil, err
	}

	return existingRestaurant, err
}
