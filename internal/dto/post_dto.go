package dto

import "time"

type CreatePostRequest struct {
	Content string `json:"content" validate:"required,min=1,max=10000" example:"This is my first post! Hello everyone!"`
}

type UpdatePostRequest struct {
	Content string `json:"content" validate:"omitempty,min=1,max=10000" example:"Updated post content"`
}

type PostResponse struct {
	ID            uint      `json:"id" example:"1"`
	Content       string    `json:"content" example:"This is my first post! Hello everyone!"`
	UserID        uint      `json:"user_id" example:"1"`
	Username      string    `json:"username" example:"john_doe"`
	CreatedAt     time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt     time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	LikesCount    int       `json:"likes_count" example:"5"`
	CommentsCount int       `json:"comments_count" example:"3"`
	IsLiked       bool      `json:"is_liked" example:"true"`
}

type PostsListResponse struct {
	Posts []PostResponse `json:"posts"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

type LikeResponse struct {
	LikesCount int  `json:"likes_count" example:"5"`
	IsLiked    bool `json:"is_liked" example:"true"`
}
