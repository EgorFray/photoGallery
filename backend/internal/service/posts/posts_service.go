package posts

import (
	"fmt"
	repository "gallery/backend/internal/repository/posts"
	"gallery/backend/internal/types"
	"gallery/backend/internal/utils"
	"mime/multipart"
	"path/filepath"
)

type PostServiceInterface interface {
	GetPosts(userId string) ([]types.PostModel, error)
	GetPostById(id int, userId string) (*types.PostDetailModel, error)
	CreatePost(file *multipart.FileHeader, description string, userId string) (*types.PostModel, error)
	SearchPosts(queryUrl, userId string) ([]types.PostModel, error)
	DeletePost(id int, userId string) error
}

type PostService struct {
	PostRepo repository.PostRepositoryInterface
}

func NewPostService(postRepository repository.PostRepositoryInterface) *PostService {
	return &PostService{PostRepo: postRepository}
}

func (s *PostService) GetPosts(userId string) ([]types.PostModel, error) {
	posts, err := s.PostRepo.DbCallGetPosts(userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetPostById(id int, userId string) (*types.PostDetailModel, error) {
	post, err := s.PostRepo.DbCallGetPostById(id, userId)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostService) CreatePost(file *multipart.FileHeader, description string, userId string) (*types.PostModel, error) {
	savePath := filepath.Join("images", "postsImg", file.Filename)
	publicPath := filepath.Join("postsImg", file.Filename)

	if err := utils.SaveFile(file, savePath); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	imagePath := "/" + publicPath
	insertedId, err := s.PostRepo.DbCallCreatePost(imagePath, description, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	post, err := s.PostRepo.DbCallGetCreatedPost(insertedId, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	return &post, nil
}

func (s *PostService) SearchPosts(queryUrl, userId string) ([]types.PostModel, error) {
	posts, err := s.PostRepo.DbCallSearchPosts(queryUrl, userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) DeletePost(id int, userId string) error {
	err := s.PostRepo.DbCallDeletePost(id, userId)
	if err != nil {
		return err
	}
	return err
}
