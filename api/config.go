package api

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func ExistingFile(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func ConfigPort() int {
	Port := viper.GetInt("Port")

	if Port == 0 {
		log.Println("No se encontro puerto, se usara el puerto 7070")
		Port = 7070
	}

	return Port
}
