package dto

import "time"

type CreateCommentRequest struct {
	Content string `json:"content" validate:"required,min=1,max=5000" example:"This is a great post!"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" validate:"omitempty,min=1,max=5000" example:"Updated comment content"`
}

type CommentResponse struct {
	ID        uint      `json:"id" example:"1"`
	Content   string    `json:"content" example:"This is a great post!"`
	UserID    uint      `json:"user_id" example:"1"`
	Username  string    `json:"username" example:"john_doe"`
	PostID    uint      `json:"post_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

type CommentsListResponse struct {
	Comments []CommentResponse `json:"comments"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
}
