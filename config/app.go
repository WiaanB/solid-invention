package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Environment string
}

func init() {
	viper.SetDefault("GOTCHA_ENV", "dev")
}

func CreateAppConfig() AppConfig {
	env := viper.GetString("GOTCHA_ENV")

	if env != "dev" && env != "prod" {
		env = "prod"
	}

	return AppConfig{
		Environment: env,
	}
}
