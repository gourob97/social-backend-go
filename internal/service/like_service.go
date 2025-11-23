package service

import (
	"social-backend/internal/apperrors"
	"social-backend/internal/dto"
	"social-backend/internal/repository/interfaces"
	serviceInterfaces "social-backend/internal/service/interfaces"
)

type likeService struct {
	likeRepo interfaces.LikeRepository
	postRepo interfaces.PostRepository
}

func NewLikeService(likeRepo interfaces.LikeRepository, postRepo interfaces.PostRepository) serviceInterfaces.LikeService {
	return &likeService{
		likeRepo: likeRepo,
		postRepo: postRepo,
	}
}

func (s *likeService) ToggleLike(userID, postID uint) (*dto.LikeResponse, error) {
	// Check if post exists
	_, err := s.postRepo.GetByID(postID)
	if err != nil {
		return nil, apperrors.ErrPostNotFound
	}
	
	// Check if user already liked the post
	isLiked, err := s.likeRepo.IsPostLikedByUser(userID, postID)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}
	
	if isLiked {
		// Unlike the post
		if err := s.likeRepo.UnlikePost(userID, postID); err != nil {
			return nil, apperrors.ErrInternalServer
		}
	} else {
		// Like the post
		if err := s.likeRepo.LikePost(userID, postID); err != nil {
			return nil, apperrors.ErrInternalServer
		}
	}
	
	// Get updated like count
	likesCount, err := s.likeRepo.GetPostLikesCount(postID)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}
	
	return &dto.LikeResponse{
		LikesCount: int(likesCount),
		IsLiked:    !isLiked, // Toggle the status
	}, nil
}

func (s *likeService) GetPostLikeStatus(userID, postID uint) (*dto.LikeResponse, error) {
	isLiked, err := s.likeRepo.IsPostLikedByUser(userID, postID)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}
	
	likesCount, err := s.likeRepo.GetPostLikesCount(postID)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}
	
	return &dto.LikeResponse{
		LikesCount: int(likesCount),
		IsLiked:    isLiked,
	}, nil
}