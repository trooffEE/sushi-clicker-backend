package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/trooffEE/sushi-clicker-backend/internal/config"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/schema"
	"log"
)

var (
	ErrConnectionFailed = errors.New("connection failed")
)

func NewDatabaseClient(cfg config.DbConfig) *sqlx.DB {
	connString := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Name, cfg.Password, cfg.Port,
	)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Panicf("Database failed to esablish connection: %w ", ErrConnectionFailed)
	}
	if err := db.Ping(); err != nil {
		log.Panicf("Database failed to be pinged: %w ", ErrConnectionFailed)
	}

	fmt.Println("Database client established connection 🥂")

	db.MustExec(schema.Schema)

	return db
}
