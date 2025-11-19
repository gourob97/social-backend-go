package repository

import (
	"social-backend/internal/apperrors"
	"social-backend/internal/model"
	"social-backend/internal/repository/interfaces"

	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) interfaces.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *model.Post) error {
	if err := r.db.Create(post).Error; err != nil {
		return apperrors.ParseDBError(err, "post")
	}
	return nil
}

func (r *postRepository) GetByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.db.Preload("User").First(&post, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrPostNotFound
		}
		return nil, apperrors.ParseDBError(err, "post")
	}
	return &post, nil
}

func (r *postRepository) GetByUserID(userID uint, limit, offset int) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	// Get total count
	if err := r.db.Model(&model.Post{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, apperrors.ParseDBError(err, "post")
	}

	// Get posts with pagination
	err := r.db.Preload("User").Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error

	if err != nil {
		return nil, 0, apperrors.ParseDBError(err, "post")
	}

	return posts, total, nil
}

func (r *postRepository) GetAll(limit, offset int) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	// Get total count
	if err := r.db.Model(&model.Post{}).Count(&total).Error; err != nil {
		return nil, 0, apperrors.ParseDBError(err, "post")
	}

	// Get posts with pagination
	err := r.db.Preload("User").
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error

	if err != nil {
		return nil, 0, apperrors.ParseDBError(err, "post")
	}

	return posts, total, nil
}

func (r *postRepository) Update(post *model.Post) error {
	if err := r.db.Save(post).Error; err != nil {
		return apperrors.ParseDBError(err, "post")
	}
	return nil
}

func (r *postRepository) Delete(id uint, userID uint) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Post{})
	if result.Error != nil {
		return apperrors.ParseDBError(result.Error, "post")
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrPostNotFound
	}
	return nil
}
