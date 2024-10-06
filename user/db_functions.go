package user

import (
	"cinnanym/model"
	"github.com/surrealdb/surrealdb.go"
)

func FindUser(DB *surrealdb.DB, email string) (model.User, error) {
	data, err := DB.Query("SELECT * FROM type::table($tb) WHERE $field = $value;", map[string]interface{}{
		"tb":    "user",
		"field": "email",
		"value": email,
	})
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	err = surrealdb.Unmarshal(data, &user)
	return user, err
}

func CreateUser(DB *surrealdb.DB, user model.User) error {
	_, err := DB.Create("user", user)
	return err
}
