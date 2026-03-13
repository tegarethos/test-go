package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	viper.AutomaticEnv()
}