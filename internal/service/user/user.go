package user

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/model"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	IsAlreadyRegistered  = errors.New("User is already registered")
	IncorrectCredentials = errors.New("Incorrect credentials")
)

type Service struct {
	usrRepo *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *Service {
	return &Service{usrRepo: ur}
}

func (s *Service) Register(email, password string) error {
	user, err := s.usrRepo.FindUserByEmail(email)
	if err != nil {
		return err
	}

	if user != nil {
		return IsAlreadyRegistered
	}
	hash, err := generatePasswordHash(password)
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

func (s *Service) Refresh(email string) (*model.User, error) {
	user, err := s.usrRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, err
}
