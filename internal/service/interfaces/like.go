package interfaces

import (
	"social-backend/internal/dto"
)

type LikeService interface {
	ToggleLike(userID, postID uint) (*dto.LikeResponse, error)
	GetPostLikeStatus(userID, postID uint) (*dto.LikeResponse, error)
}