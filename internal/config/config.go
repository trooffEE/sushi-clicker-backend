package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var IsDevelopment = os.Getenv("IN_CONTAINER") == ""

func init() {
	if IsDevelopment {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
