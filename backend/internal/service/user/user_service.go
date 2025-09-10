package user

import (
	repository "gallery/backend/internal/repository/user"
	"gallery/backend/internal/types"
)

type UserServiceInterface interface {
	CreateUser(req types.UserRequest, hashedPassword string) (*int, error)
}

type UserService struct {
	UserRepo repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{UserRepo: userRepository}
}

func (s *UserService) CreateUser(req types.UserRequest, hashedPassword string) (*int, error) {
	userId, err := s.UserRepo.DbCallCreateUser(req.Name, req.Email, hashedPassword, req.Avatar)
	if err != nil {
		return nil, err
	}
	return &userId, nil
}