package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/model"
	"go.uber.org/zap"
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
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
	_, err := r.db.NamedExec("INSERT INTO users (email, hash, token_sugar) VALUES (:email, :hash, :token_sugar)", user)
	if err != nil {
		zap.L().Error(err.Error(), zap.String("email", user.Email))
		return err
	}

	return nil
}
