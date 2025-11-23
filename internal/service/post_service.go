package service

import (
	"social-backend/internal/apperrors"
	"social-backend/internal/dto"
	"social-backend/internal/model"
	"social-backend/internal/repository/interfaces"
	serviceInterfaces "social-backend/internal/service/interfaces"
)

type postService struct {
	postRepo    interfaces.PostRepository
	likeRepo    interfaces.LikeRepository
	commentRepo interfaces.CommentRepository
}

func NewPostService(postRepo interfaces.PostRepository, likeRepo interfaces.LikeRepository, commentRepo interfaces.CommentRepository) serviceInterfaces.PostService {
	return &postService{
		postRepo:    postRepo,
		likeRepo:    likeRepo,
		commentRepo: commentRepo,
	}
}

func (s *postService) CreatePost(userID uint, req *dto.CreatePostRequest) (*dto.PostResponse, error) {
	post := &model.Post{
		Content: req.Content,
		UserID:  userID,
	}

	if err := s.postRepo.Create(post); err != nil {
		return nil, err
	}

	// Fetch the created post with user data
	createdPost, err := s.postRepo.GetByID(post.ID)
	if err != nil {
		return nil, err
	}

	return s.convertToPostResponse(createdPost), nil
}

func (s *postService) GetPost(id uint) (*dto.PostResponse, error) {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.convertToPostResponse(post), nil
}

func (s *postService) GetUserPosts(userID uint, page, limit int) (*dto.PostsListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit
	posts, total, err := s.postRepo.GetByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	postResponses := make([]dto.PostResponse, 0, len(posts))
	for _, post := range posts {
		postResponses = append(postResponses, *s.convertToPostResponse(&post))
	}

	return &dto.PostsListResponse{
		Posts: postResponses,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *postService) GetAllPosts(page, limit int) (*dto.PostsListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit
	posts, total, err := s.postRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	postResponses := make([]dto.PostResponse, 0, len(posts))
	for _, post := range posts {
		postResponses = append(postResponses, *s.convertToPostResponse(&post))
	}

	return &dto.PostsListResponse{
		Posts: postResponses,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *postService) UpdatePost(id uint, userID uint, req *dto.UpdatePostRequest) (*dto.PostResponse, error) {
	// Get existing post
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check ownership
	if post.UserID != userID {
		return nil, apperrors.ErrUnauthorized
	}

	// Update fields if provided
	if req.Content != "" {
		post.Content = req.Content
	}

	if err := s.postRepo.Update(post); err != nil {
		return nil, err
	}

	// Fetch updated post
	updatedPost, err := s.postRepo.GetByID(post.ID)
	if err != nil {
		return nil, err
	}

	return s.convertToPostResponse(updatedPost), nil
}

func (s *postService) DeletePost(id uint, userID uint) error {
	return s.postRepo.Delete(id, userID)
}

func (s *postService) convertToPostResponse(post *model.Post) *dto.PostResponse {
	// Get likes count
	likesCount, _ := s.likeRepo.GetPostLikesCount(post.ID)

	// Get comments count
	commentsCount, _ := s.commentRepo.GetCommentsCountByPostID(post.ID)

	return &dto.PostResponse{
		ID:            post.ID,
		Content:       post.Content,
		UserID:        post.UserID,
		Username:      post.User.Username,
		CreatedAt:     post.CreatedAt,
		UpdatedAt:     post.UpdatedAt,
		LikesCount:    int(likesCount),
		CommentsCount: int(commentsCount),
		IsLiked:       false, // TODO: Check if current user liked this post
	}
}
