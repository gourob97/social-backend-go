package interfaces

import (
	"social-backend/internal/model"
)

type CommentRepository interface {
	Create(comment *model.Comment) (*model.Comment, error)
	GetByID(id uint) (*model.Comment, error)
	GetByPostID(postID uint, page, limit int) ([]model.Comment, int64, error)
	Update(comment *model.Comment) (*model.Comment, error)
	Delete(id uint) error
	GetCommentsCountByPostID(postID uint) (int64, error)
	GetByUserID(userID uint, page, limit int) ([]model.Comment, int64, error)
}
