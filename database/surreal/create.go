package surreal

import (
	"cinnanym/maps"
	"fmt"
)

type CreatePayload struct {
	Table      string
	Identifier string
	Data       maps.Map
}

type GenericModel interface {
	VerifyCreate() error
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

func Create[GM GenericModel](payload CreatePayload, model GM) error {
	err := model.VerifyCreate()
	if err != nil {
		return err
	}

	_, err = DB.Create(payload.ID(), payload.Data)
	return err
}
