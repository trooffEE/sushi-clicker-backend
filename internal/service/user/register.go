package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/model"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
)

func (s *Service) Register(email, password string) error {
	user, err := s.usrRepo.FindUserByEmail(email)
	if err != nil {
		return err
	}

	if user != nil {
		return IsAlreadyRegistered
	}
	hash, err := lib.GeneratePasswordHash(password)
	if err != nil {
		fmt.Printf("Something went wrong on registration\n")
		return err
	}

	usr := model.User{
		Email: email,
		Hash:  hash,
		Sugar: uuid.NewString(),
	}
	err = s.usrRepo.CreateUser(&usr)
	if err != nil {
		return err
	}

	return nil
}
