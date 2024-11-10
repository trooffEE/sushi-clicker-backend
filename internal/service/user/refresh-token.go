package user

import "github.com/trooffEE/sushi-clicker-backend/internal/db/model"

func (s *Service) RefreshToken(email string) (*model.User, error) {
	user, err := s.usrRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, err
}
