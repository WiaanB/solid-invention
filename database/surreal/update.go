package surreal

type UpdatePayload struct {
	ID   string
	Data map[string]interface{}
}

func Update(payload UpdatePayload) error {
	_, err := DB.Update(payload.ID, payload.Data)
	return err
}
