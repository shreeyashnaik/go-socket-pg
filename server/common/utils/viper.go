package utils

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadEnv() {
	currentPath, _ := os.Getwd()
	log.Printf("Current path: %s", currentPath)

	envPath := currentPath + "/.env"
	viper.SetConfigFile(envPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
