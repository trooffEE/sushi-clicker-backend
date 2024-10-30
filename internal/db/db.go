package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
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
}

func NewDB(args ...func(*Config)) *sqlx.DB {
	cfg := Config{
		host:     "localhost", // TODO
		user:     os.Getenv("DB_USER"),
		dbname:   os.Getenv("DB_NAME"),
		password: os.Getenv("DB_PASSWORD"),
	}

	connString := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.host, cfg.user, cfg.dbname, cfg.password,
	)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Panicf("Establish failed: %w ", ErrConnectionFailed)
	}

	//TODO
	db.MustExec(schema.Schema)

	return db
}
