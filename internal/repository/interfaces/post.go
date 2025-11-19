package interfaces

import "social-backend/internal/model"

type PostRepository interface {
	Create(post *model.Post) error
	GetByID(id uint) (*model.Post, error)
	GetByUserID(userID uint, limit, offset int) ([]model.Post, int64, error)
	GetAll(limit, offset int) ([]model.Post, int64, error)
	Update(post *model.Post) error
	Delete(id uint, userID uint) error
}
