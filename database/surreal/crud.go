package surreal

import (
	"cinnanym/maps"
	"encoding/json"
	"fmt"
)

type CreatePayload struct {
	Table      string
	Identifier string
	Data       maps.Map
}

func (p *CreatePayload) ID() string {

	if p.Identifier == "" {
		return fmt.Sprintf("%s", p.Table)
	}

	value, ok := p.Data[p.Identifier]
	if ok {
		return fmt.Sprintf("%s:%s", p.Table, value)
	}

	return fmt.Sprintf("%s", p.Table)
}

type UpdatePayload struct {
	ID   string
	Data maps.Map
}

type GenericModel interface {
	VerifyCreate() error
}

func Create[GM GenericModel](payload CreatePayload, model GM) error {
	err := model.VerifyCreate()
	if err != nil {
		return err
	}

	_, err = DB.Create(payload.ID(), payload.Data)
	return err
}

func Delete(id string) error {
	_, err := DB.Delete(id)
	return err
}

func Update(payload UpdatePayload) error {
	entity, err := FindOne(payload.ID)
	if err != nil {
		return err
	}

	if entity == nil {
		return fmt.Errorf("entity not found")
	}

	_, err = DB.Update(payload.ID, payload.Data)
	return err
}

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
