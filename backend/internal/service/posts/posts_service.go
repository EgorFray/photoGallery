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
	GetPosts() ([]types.PostModel, error)
	GetPostById(id int) (*types.PostDetailModel, error)
	CreatePost(file *multipart.FileHeader, description string) (*types.PostModel, error)
	SearchPosts(queryUrl string) ([]types.PostModel, error)
	DeletePost(id int) error
}

type PostService struct {
	PostRepo repository.PostRepositoryInterface
}

func NewPostService(postRepository repository.PostRepositoryInterface) *PostService {
	return &PostService{PostRepo: postRepository}
}

func (s *PostService) GetPosts() ([]types.PostModel, error) {
	posts, err := s.PostRepo.DbCallGetPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetPostById(id int) (*types.PostDetailModel, error) {
	post, err := s.PostRepo.DbCallGetPostById(id)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostService) CreatePost(file *multipart.FileHeader, description string) (*types.PostModel, error) {
	filePath := filepath.Join("postsImg", file.Filename)

	if err := utils.SaveFile(file, filePath); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	imagePath := "/" + filePath
	insertedId, err := s.PostRepo.DbCallCreatePost(imagePath, description)
	if err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	post, err := s.PostRepo.DbCallGetCreatedPost(insertedId)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	return &post, nil
}

func (s *PostService) SearchPosts(queryUrl string) ([]types.PostModel, error) {
	posts, err := s.PostRepo.DbCallSearchPosts(queryUrl)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) DeletePost(id int) error {
	err := s.PostRepo.DbCallDeletePost(id)
	if err != nil {
		return err
	}
	return err
}
