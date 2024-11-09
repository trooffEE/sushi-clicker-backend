package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/model"
	"log"
)

type UserRepositoryInterface interface {
	FindUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * from users WHERE email = $1", email)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
	_, err := r.db.NamedExec("INSERT INTO users (email, hash) VALUES (:email, :hash)", user)
	if err != nil {
		log.Printf("%w\n", err)
		return err
	}

	return nil
}
