package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Environment string
}

func init() {
	viper.SetDefault("ENVIRONMENT", "dev")
}

func CreateAppConfig() AppConfig {
	env := viper.GetString("ENVIRONMENT")

	if env != "dev" && env != "prod" {
		env = "prod"
	}

	return AppConfig{
		Environment: env,
	}
}
