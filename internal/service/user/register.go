package user

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/model"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
	"go.uber.org/zap"
)

func (s *Service) Register(email, password string) (*model.User, error) {
	user, err := s.usrRepo.FindUserByEmail(email)
	if !errors.Is(sql.ErrNoRows, err) && err != nil {
		return nil, err
	}

	if user != nil {
		return nil, IsAlreadyRegistered
	}
	hash, err := lib.GeneratePasswordHash(password)
	if err != nil {
		zap.L().Error("Failed to generate password hash", zap.Error(err))
		return nil, err
	}

	usr := model.User{
		Email: email,
		Hash:  hash,
		Sugar: uuid.NewString(),
	}
	err = s.usrRepo.CreateUser(&usr)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}
