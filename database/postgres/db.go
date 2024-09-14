package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotcha/config"
	"gotcha/model"
)

var Database *gorm.DB

func InitialiseDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", "localhost", config.Get().DB.User, config.Get().DB.Pass, config.Get().DB.Database)
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = Database.AutoMigrate(&model.User{})
	if err != nil {
		panic("Failed to connect to database!")
	}
}
