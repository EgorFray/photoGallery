package user

import (
	"fmt"
	repository "gallery/backend/internal/repository/user"
	"gallery/backend/internal/types"
	"gallery/backend/internal/utils"
	"mime/multipart"
	"path/filepath"
)

type UserServiceInterface interface {
	CreateUser(req types.UserRequest, hashedPassword string, file *multipart.FileHeader) (*int, error)
	GetUserByEmail(email string) (*types.UserModel, error)
	UpdateUser(id, name, password string, file *multipart.FileHeader) error
}

type UserService struct {
	UserRepo repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{UserRepo: userRepository}
}

func (s *UserService) CreateUser(req types.UserRequest, hashedPassword string, file *multipart.FileHeader) (*int, error) {
	savePath := filepath.Join("images", "avatars", file.Filename)
	publicPath := filepath.Join("avatars", file.Filename)

	if err := utils.SaveFile(file, savePath); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	imagePath := "/" + publicPath
	userId, err := s.UserRepo.DbCallCreateUser(req.Name, req.Email, hashedPassword, imagePath)
	if err != nil {
		return nil, err
	}
	return &userId, nil
}

func (s *UserService) GetUserByEmail(email string) (*types.UserModel, error) {
	userData, err := s.UserRepo.DbCallGetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

func (s *UserService) UpdateUser(id, name, password string, file *multipart.FileHeader) error {
	savePath := filepath.Join("images", "avatars", file.Filename)
	publicPath := filepath.Join("avatars", file.Filename)

	if err := utils.SaveFile(file, savePath); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	imagePath := "/" + publicPath

	if password != "" {
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return err
		}
	}
	
	err := s.UserRepo.DbCallUpdateUser(id, name, hashedPassword, imagePath)
	if err != nil {
		return err
	}
	return nil
}