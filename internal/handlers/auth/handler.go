package auth

import "github.com/trooffEE/sushi-clicker-backend/internal/service/user"

type Handler struct {
	UserService *user.Service
}

func NewHandler(userService *user.Service) *Handler {
	return &Handler{}
}
