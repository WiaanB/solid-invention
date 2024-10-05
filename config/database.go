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
	viper.SetDefault("DB_USER", "admin")
	viper.SetDefault("DB_PASS", "admin")
	viper.SetDefault("DB_NAMESPACE", "my_namespace")
	viper.SetDefault("DB_DATABASE", "my_database")
}

func CreateDBConfig() DBConfig {
	return DBConfig{
		User:      viper.GetString("DB_USER"),
		Pass:      viper.GetString("DB_PASS"),
		Namespace: viper.GetString("DB_NAMESPACE"),
		Database:  viper.GetString("DB_DATABASE"),
	}
}
