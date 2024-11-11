package user

import (
	"github.com/trooffEE/sushi-clicker-backend/internal/db/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(email, password string) (*model.User, error) {
	user, err := s.usrRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	if err != nil {
		return nil, IncorrectCredentials
	}

	return user, err
}