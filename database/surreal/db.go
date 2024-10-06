package surreal

import (
	"cinnanym/config"
	"github.com/surrealdb/surrealdb.go"
)

var DB *surrealdb.DB

func SetupDB() {
	var err error
	DB, err = surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		panic(err)
	}

	if _, err = DB.Signin(map[string]interface{}{
		"user": config.Get().DB.User,
		"pass": config.Get().DB.Pass,
	}); err != nil {
		panic(err)
	}

	if _, err = DB.Use(config.Get().DB.Namespace, config.Get().DB.Database); err != nil {
		panic(err)
	}
}
