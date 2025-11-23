package repository

import (
	"social-backend/internal/model"
	"social-backend/internal/repository/interfaces"
	"gorm.io/gorm"
)

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) interfaces.LikeRepository {
	return &likeRepository{db: db}
}

func (r *likeRepository) LikePost(userID, postID uint) error {
	like := model.Like{
		UserID: userID,
		PostID: postID,
	}
	
	// Use FirstOrCreate to prevent duplicate likes
	result := r.db.Where("user_id = ? AND post_id = ?", userID, postID).FirstOrCreate(&like)
	return result.Error
}

func (r *likeRepository) UnlikePost(userID, postID uint) error {
	return r.db.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&model.Like{}).Error
}

func (r *likeRepository) IsPostLikedByUser(userID, postID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Like{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count).Error
	return count > 0, err
}

func (r *likeRepository) GetPostLikesCount(postID uint) (int64, error) {
	var count int64
	err := r.db.Model(&model.Like{}).Where("post_id = ?", postID).Count(&count).Error
	return count, err
}

func (r *likeRepository) GetLikesByUser(userID uint) ([]model.Like, error) {
	var likes []model.Like
	err := r.db.Where("user_id = ?", userID).Preload("Post").Find(&likes).Error
	return likes, err
}