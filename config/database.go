package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	User      string
	Pass      string
	Namespace string
	Database  string
}

func init() {
	viper.SetDefault("GOTCHA_DB_USER", "admin")
	viper.SetDefault("GOTCHA_DB_PASS", "admin")
	viper.SetDefault("GOTCHA_DB_NAMESPACE", "gotcha")
	viper.SetDefault("GOTCHA_DB_DATABASE", "gotcha")
}

func CreateDBConfig() DBConfig {
	return DBConfig{
		User:      viper.GetString("GOTCHA_DB_USER"),
		Pass:      viper.GetString("GOTCHA_DB_PASS"),
		Namespace: viper.GetString("GOTCHA_DB_NAMESPACE"),
		Database:  viper.GetString("GOTCHA_DB_DATABASE"),
	}
}
