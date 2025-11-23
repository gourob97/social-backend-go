package interfaces

import (
	"social-backend/internal/model"
)

type LikeRepository interface {
	LikePost(userID, postID uint) error
	UnlikePost(userID, postID uint) error
	IsPostLikedByUser(userID, postID uint) (bool, error)
	GetPostLikesCount(postID uint) (int64, error)
	GetLikesByUser(userID uint) ([]model.Like, error)
}