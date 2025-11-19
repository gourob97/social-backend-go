package interfaces

import "social-backend/internal/dto"

type PostService interface {
	CreatePost(userID uint, req *dto.CreatePostRequest) (*dto.PostResponse, error)
	GetPost(id uint) (*dto.PostResponse, error)
	GetUserPosts(userID uint, page, limit int) (*dto.PostsListResponse, error)
	GetAllPosts(page, limit int) (*dto.PostsListResponse, error)
	UpdatePost(id uint, userID uint, req *dto.UpdatePostRequest) (*dto.PostResponse, error)
	DeletePost(id uint, userID uint) error
}
