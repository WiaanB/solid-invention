package config

import "github.com/spf13/viper"

type ApiConfig struct {
	Port int
}

func init() {
	viper.SetDefault("GOTCHA_API_PORT", 8080)
}

func CreateAPIConfig() ApiConfig {
	return ApiConfig{
		Port: viper.GetInt("GOTCHA_API_PORT"),
	}
}
