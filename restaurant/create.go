package restaurant

import (
	"cinnanym/database/surreal"
	"cinnanym/model"
	"errors"
)

func Create(restaurant model.Restaurant) (model.Restaurant, error) {
	// ensure the required fields are there
	if restaurant.Name == "" || restaurant.Address == "" || restaurant.Type == "" {
		return model.Restaurant{}, errors.New("required fields are missing, try again")
	}

	err := surreal.Create(surreal.CreatePayload{
		Table:      "restaurant",
		Identifier: "",
		Data:       restaurant.ToMap(),
	})
	if err != nil {
		return model.Restaurant{}, err
	}

	return restaurant, nil
}
