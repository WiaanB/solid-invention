package surreal

import (
	"cinnanym/maps"
	"fmt"
)

type UpdatePayload struct {
	ID   string
	Data maps.Map
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
