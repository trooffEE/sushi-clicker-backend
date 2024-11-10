package user

import (
	"github.com/trooffEE/sushi-clicker-backend/internal/db/repository"
)

type Service struct {
	usrRepo *repository.UserRepository
}

// TODO Refactor https://youtu.be/hDwqFRUuykQ?t=868
func NewUserService(ur *repository.UserRepository) *Service {
	return &Service{usrRepo: ur}
}
