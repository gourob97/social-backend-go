package dto

import "time"

type CreatePostRequest struct {
	Content string `json:"content" validate:"required,min=1,max=10000"`
}

type UpdatePostRequest struct {
	Content string `json:"content" validate:"omitempty,min=1,max=10000"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostsListResponse struct {
	Posts []PostResponse `json:"posts"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}
