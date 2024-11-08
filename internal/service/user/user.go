package user

import (
	"errors"
	"fmt"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/model"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/repository"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
)

var (
	IsAlreadyRegistered = errors.New("user already registered")
)

type Service struct {
	usrRepo repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *Service {
	return &Service{usrRepo: *ur}
}

func (s *Service) Register(email, password string) error {
	user, err := s.usrRepo.FindUserByEmail(email)
	if err != nil {
		fmt.Printf("Something went wrong on registration user with email %s\n", email)
		return err
	}

	if user != nil {
		return IsAlreadyRegistered
	}
	hash, err := lib.GenerateJWTToken(password)
	if err != nil {
		fmt.Printf("Something went wrong on registration\n")
		return err
	}

	usr := model.User{
		Email: email,
		Hash:  hash,
	}
	err = s.usrRepo.CreateUser(&usr)
	if err != nil {
		fmt.Printf("Something went wrong on registration #%v \n", usr)
	}

	return nil
}
