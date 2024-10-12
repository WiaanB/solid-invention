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
	fmt.Println("Updating", payload.ID, payload.Data)
	_, err := DB.Update(payload.ID, payload.Data)
	return err
}
