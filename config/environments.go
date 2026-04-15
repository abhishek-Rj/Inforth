package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MainDir		string 	`envconfig:"MAIN_DIRECTORY" required:"true"`
}

var App Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = envconfig.Process("", &App)
	if err != nil {
		log.Fatal(err)
	}
}