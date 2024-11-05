package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/model"
	"log"
)

type UserRepositoryInterface interface {
	Register(user model.User) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Register(user model.User) error {
	_, err := r.db.NamedExec(
		`INSERT INTO users (email, hash) VALUES (:email, :hash)`,
		user)

	if err != nil {
		log.Printf("%w\n", err)
		return err
	}

	return nil
}
