package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/schema"
	"log"
	"os"
)

var (
	ErrConnectionFailed = errors.New("connection failed")
)

type Config struct {
	host     string
	user     string
	dbname   string
	password string
	port     string
}

func NewDatabaseClient() *sqlx.DB {
	err := godotenv.Load()
	cfg := Config{
		host:     "sushi_db_container",
		user:     os.Getenv("DB_USER"),
		dbname:   os.Getenv("DB_NAME"),
		password: os.Getenv("DB_PASSWORD"),
		port:     os.Getenv("DB_PORT"),
	}

	connString := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.host, cfg.user, cfg.dbname, cfg.password, cfg.port,
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
