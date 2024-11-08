package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if os.Getenv("IN_CONTAINER") == "" {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
