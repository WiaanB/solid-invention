package user

import (
	"cinnanym/database/surreal"
)

func Delete(id string) error {
	return surreal.Delete(id)
}
