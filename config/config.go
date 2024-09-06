package config

import (
	"github.com/spf13/viper"
	"log"
)

type GotchaConfig struct {
	App AppConfig
	Api ApiConfig
}

func Get() *GotchaConfig {
	return Config
}

var Config *GotchaConfig

func init() {
	// read the .env file
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	Config = &GotchaConfig{
		App: CreateAppConfig(),
		Api: CreateAPIConfig(),
	}
}
