package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/trooffEE/sushi-clicker-backend/internal/config"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/schema"
	"go.uber.org/zap"
)

func NewDatabaseClient(cfg config.DbConfig) *sqlx.DB {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		zap.L().Panic("Database failed to establish connection: ", zap.Error(err))
	}
	if err := db.Ping(); err != nil {
		zap.L().Panic("Database failed to be pinged: ", zap.Error(err))
	}

	m, err := migrate.New(
		"file://internal/db/migrations",
		connString,
	)
	if err != nil {
		zap.L().Fatal("Database migration failed: ", zap.Error(err))
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		zap.L().Fatal("Failed to run migrations: ", zap.Error(err))
	}

	zap.L().Info("Database connection established 🥂")

	db.MustExec(schema.Schema)

	return db
}
