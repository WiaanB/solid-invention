package surreal

import (
	"fmt"
)

type CreatePayload struct {
	Table      string
	Identifier string
	Data       map[string]interface{}
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

func Create(payload CreatePayload) error {
	_, err := DB.Create(payload.ID(), payload.Data)
	return err
}
