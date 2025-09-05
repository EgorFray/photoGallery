package posts

import repository "gallery/backend/internal/repository/posts"

type PostService struct {
	PostRepo repository.PostRepositoryInterface
}

func NewPostService(postRepository repository.PostRepositoryInterface) *PostService {
	return &PostService{PostRepo: postRepository}
}

