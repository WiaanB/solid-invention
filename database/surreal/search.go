package surreal

import (
	"encoding/json"
)

func FindOne[T any](id string) (T, error) {
	var result T

	data, err := DB.Select(id)
	if err != nil {
		return result, err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
