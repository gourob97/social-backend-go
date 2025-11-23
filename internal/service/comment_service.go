package service

import (
	"social-backend/internal/apperrors"
	"social-backend/internal/dto"
	"social-backend/internal/model"
	"social-backend/internal/repository/interfaces"
	serviceInterfaces "social-backend/internal/service/interfaces"
)

type commentService struct {
	commentRepo interfaces.CommentRepository
	postRepo    interfaces.PostRepository
}

func NewCommentService(commentRepo interfaces.CommentRepository, postRepo interfaces.PostRepository) serviceInterfaces.CommentService {
	return &commentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

func (s *commentService) CreateComment(userID uint, postID uint, req *dto.CreateCommentRequest) (*dto.CommentResponse, error) {
	// Check if post exists
	_, err := s.postRepo.GetByID(postID)
	if err != nil {
		return nil, apperrors.ErrPostNotFound
	}

	comment := &model.Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  postID,
	}

	createdComment, err := s.commentRepo.Create(comment)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}

	return &dto.CommentResponse{
		ID:        createdComment.ID,
		Content:   createdComment.Content,
		UserID:    createdComment.UserID,
		Username:  createdComment.User.Username,
		PostID:    createdComment.PostID,
		CreatedAt: createdComment.CreatedAt,
		UpdatedAt: createdComment.UpdatedAt,
	}, nil
}

func (s *commentService) GetComment(id uint) (*dto.CommentResponse, error) {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return nil, apperrors.ErrPostNotFound
	}

	return &dto.CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		UserID:    comment.UserID,
		Username:  comment.User.Username,
		PostID:    comment.PostID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}, nil
}

func (s *commentService) GetCommentsByPostID(postID uint, page, limit int) (*dto.CommentsListResponse, error) {
	// Check if post exists
	_, err := s.postRepo.GetByID(postID)
	if err != nil {
		return nil, apperrors.ErrPostNotFound
	}

	comments, total, err := s.commentRepo.GetByPostID(postID, page, limit)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}

	commentResponses := make([]dto.CommentResponse, len(comments))
	for i, comment := range comments {
		commentResponses[i] = dto.CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			UserID:    comment.UserID,
			Username:  comment.User.Username,
			PostID:    comment.PostID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
	}

	return &dto.CommentsListResponse{
		Comments: commentResponses,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

func (s *commentService) UpdateComment(userID, commentID uint, req *dto.UpdateCommentRequest) (*dto.CommentResponse, error) {
	comment, err := s.commentRepo.GetByID(commentID)
	if err != nil {
		return nil, apperrors.ErrPostNotFound
	}

	// Check if user owns the comment
	if comment.UserID != userID {
		return nil, apperrors.ErrForbidden
	}

	comment.Content = req.Content

	updatedComment, err := s.commentRepo.Update(comment)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}

	return &dto.CommentResponse{
		ID:        updatedComment.ID,
		Content:   updatedComment.Content,
		UserID:    updatedComment.UserID,
		Username:  updatedComment.User.Username,
		PostID:    updatedComment.PostID,
		CreatedAt: updatedComment.CreatedAt,
		UpdatedAt: updatedComment.UpdatedAt,
	}, nil
}

func (s *commentService) DeleteComment(userID, commentID uint) error {
	comment, err := s.commentRepo.GetByID(commentID)
	if err != nil {
		return apperrors.ErrPostNotFound
	}

	// Check if user owns the comment
	if comment.UserID != userID {
		return apperrors.ErrForbidden
	}

	if err := s.commentRepo.Delete(commentID); err != nil {
		return apperrors.ErrInternalServer
	}

	return nil
}

func (s *commentService) GetCommentsByUserID(userID uint, page, limit int) (*dto.CommentsListResponse, error) {
	comments, total, err := s.commentRepo.GetByUserID(userID, page, limit)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}

	commentResponses := make([]dto.CommentResponse, len(comments))
	for i, comment := range comments {
		commentResponses[i] = dto.CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			UserID:    comment.UserID,
			Username:  comment.User.Username,
			PostID:    comment.PostID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
	}

	return &dto.CommentsListResponse{
		Comments: commentResponses,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}
