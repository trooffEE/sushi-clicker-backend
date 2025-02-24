package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/trooffEE/sushi-clicker-backend/internal/config"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/schema"
	"go.uber.org/zap"
)

func NewDatabaseClient(cfg config.DbConfig) *sqlx.DB {
	connString := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Name, cfg.Password, cfg.Port,
	)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		zap.L().Panic("Database failed to establish connection: ", zap.Error(err))
	}
	if err := db.Ping(); err != nil {
		zap.L().Panic("Database failed to be pinged: ", zap.Error(err))
	}

	zap.L().Info("Database connection established 🥂")

	db.MustExec(schema.Schema)

	return db
}
