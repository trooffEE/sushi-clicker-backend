package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type ApplicationConfig struct {
	Database DbConfig
	Jwt      JwtConfig
}

type DbConfig struct {
	Host     string
	Name     string
	User     string
	Port     string
	Password string
}

type JwtConfig struct {
	Secret        string
	SecretRefresh string
}

// IsDevelopment TODO remove wierd logic for in container
var IsDevelopment = os.Getenv("IN_CONTAINER") == ""

func NewApplicationConfig() ApplicationConfig {
	if IsDevelopment {
		if err := godotenv.Load(".env.local"); err != nil {
			zap.L().Fatal("Error loading .env file", zap.Error(err))
		}
	} else {
		if err := godotenv.Load(); err != nil {
			zap.L().Fatal("Error loading .env file", zap.Error(err))
		}
	}

	return ApplicationConfig{
		Database: DbConfig{
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Port:     os.Getenv("DB_PORT"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		Jwt: JwtConfig{
			Secret:        os.Getenv("JWT_SECRET_KEY"),
			SecretRefresh: os.Getenv("JWT_SECRET_KEY_REFRESH"),
		},
	}

}
