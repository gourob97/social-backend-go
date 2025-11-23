package interfaces

import (
	"social-backend/internal/dto"
)

type CommentService interface {
	CreateComment(userID uint, postID uint, req *dto.CreateCommentRequest) (*dto.CommentResponse, error)
	GetComment(id uint) (*dto.CommentResponse, error)
	GetCommentsByPostID(postID uint, page, limit int) (*dto.CommentsListResponse, error)
	UpdateComment(userID, commentID uint, req *dto.UpdateCommentRequest) (*dto.CommentResponse, error)
	DeleteComment(userID, commentID uint) error
	GetCommentsByUserID(userID uint, page, limit int) (*dto.CommentsListResponse, error)
}
