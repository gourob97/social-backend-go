package repository

import (
	"social-backend/internal/apperrors"
	"social-backend/internal/model"
	"social-backend/internal/repository/interfaces"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) interfaces.CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *model.Comment) (*model.Comment, error) {
	if err := r.db.Create(comment).Error; err != nil {
		return nil, apperrors.ParseDBError(err, "comment")
	}

	// Fetch the created comment with user data
	if err := r.db.Preload("User").First(comment, comment.ID).Error; err != nil {
		return nil, apperrors.ParseDBError(err, "comment")
	}

	return comment, nil
}

func (r *commentRepository) GetByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	if err := r.db.Preload("User").Preload("Post").First(&comment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrPostNotFound
		}
		return nil, apperrors.ParseDBError(err, "comment")
	}
	return &comment, nil
}

func (r *commentRepository) GetByPostID(postID uint, page, limit int) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Count total comments for this post
	if err := r.db.Model(&model.Comment{}).Where("post_id = ?", postID).Count(&total).Error; err != nil {
		return nil, 0, apperrors.ParseDBError(err, "comment")
	}

	// Get comments with pagination
	err := r.db.Where("post_id = ?", postID).
		Preload("User").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&comments).Error

	if err != nil {
		return nil, 0, apperrors.ParseDBError(err, "comment")
	}

	return comments, total, nil
}

func (r *commentRepository) Update(comment *model.Comment) (*model.Comment, error) {
	if err := r.db.Save(comment).Error; err != nil {
		return nil, apperrors.ParseDBError(err, "comment")
	}

	// Fetch updated comment with user data
	if err := r.db.Preload("User").First(comment, comment.ID).Error; err != nil {
		return nil, apperrors.ParseDBError(err, "comment")
	}

	return comment, nil
}

func (r *commentRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Comment{}, id)
	if result.Error != nil {
		return apperrors.ParseDBError(result.Error, "comment")
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrPostNotFound
	}
	return nil
}

func (r *commentRepository) GetCommentsCountByPostID(postID uint) (int64, error) {
	var count int64
	err := r.db.Model(&model.Comment{}).Where("post_id = ?", postID).Count(&count).Error
	if err != nil {
		return 0, apperrors.ParseDBError(err, "comment")
	}
	return count, nil
}

func (r *commentRepository) GetByUserID(userID uint, page, limit int) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Count total comments by user
	if err := r.db.Model(&model.Comment{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, apperrors.ParseDBError(err, "comment")
	}

	// Get comments with pagination
	err := r.db.Where("user_id = ?", userID).
		Preload("User").
		Preload("Post").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&comments).Error

	if err != nil {
		return nil, 0, apperrors.ParseDBError(err, "comment")
	}

	return comments, total, nil
}
