package surreal

import (
	"cinnanym/maps"
)

type UpdatePayload struct {
	ID   string
	Data maps.Map
}

func Update(payload UpdatePayload) error {
	_, err := DB.Update(payload.ID, payload.Data)
	return err
}
