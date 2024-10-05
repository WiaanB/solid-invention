package config

import (
	"github.com/spf13/viper"
	"log"
)

type CinnanymConfig struct {
	App AppConfig
	Api ApiConfig
	DB  DBConfig
}

func Get() *CinnanymConfig {
	return Config
}

var Config *CinnanymConfig

func init() {
	// read the .env file
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	Config = &CinnanymConfig{
		App: CreateAppConfig(),
		Api: CreateAPIConfig(),
		DB:  CreateDBConfig(),
	}
}
