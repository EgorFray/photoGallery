package posts

import (
	repository "gallery/backend/internal/repository/posts"
	"gallery/backend/internal/types"
)

type PostServiceInterface interface {
	GetPosts() ([]types.PostModel, error)
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

