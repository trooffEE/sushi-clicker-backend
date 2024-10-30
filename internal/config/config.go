package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Cfg struct{}

func initConfig() *Cfg {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[error] Error loading .env file")
	}

	return &Cfg{}
}

var Config = initConfig()
