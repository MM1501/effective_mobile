package config

import (
	"log"

	"github.com/spf13/viper"
)

var DBConnectionString string

func InitConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}
	DBConnectionString = viper.GetString("DB_CONNECTION_STRING")
}
