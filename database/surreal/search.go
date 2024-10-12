package surreal

import (
	"cinnanym/maps"
	"encoding/json"
	"fmt"
)

func FindOne(id string) (maps.Map, error) {
	data, err := DB.Select(id)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var result maps.Map
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type FindAllPayload struct {
	Table string
	Size  int
	Page  int
}

func FindAll(payload FindAllPayload) (map[string]interface{}, error) {
	data, err := DB.Query(fmt.Sprintf("SELECT * FROM %s LIMIT %d START %d;", payload.Table, payload.Size, payload.Size*payload.Page), map[string]interface{}{})

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result[0], nil
}
