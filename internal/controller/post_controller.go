package controller

import (
	"social-backend/internal/apperrors"
	"social-backend/internal/dto"
	"social-backend/internal/response"
	"social-backend/internal/service/interfaces"
	"social-backend/internal/validation"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	postService interfaces.PostService
}

func NewPostController(postService interfaces.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

// CreatePostController handles creating a new post
// @Summary Create a new post
// @Description Create a new post with content
// @Tags posts
// @Accept json
// @Produce json
// @Param post body dto.CreatePostRequest true "Post data"
// @Success 201 {object} response.APIResponse{data=dto.PostResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Security ApiKeyAuth
// @Router /api/posts [post]
func (pc *PostController) CreatePostController(c echo.Context) error {
	var req dto.CreatePostRequest

	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, apperrors.ErrInvalidInput)
	}

	if err := validation.ValidateStruct(&req); err != nil {
		return response.ErrorResponse(c, err)
	}

	// Get user ID from JWT token (this will be set by auth middleware)
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return response.ErrorResponse(c, apperrors.ErrUnauthorized)
	}

	post, err := pc.postService.CreatePost(userID, &req)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, 201, "Post created successfully", post)
}

// GetPostController handles retrieving a specific post
// @Summary Get a post by ID
// @Description Retrieve a specific post by its ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} response.APIResponse{data=dto.PostResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/posts/{id} [get]
func (pc *PostController) GetPostController(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.ErrorResponse(c, apperrors.ErrInvalidInput)
	}

	post, err := pc.postService.GetPost(uint(id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, 200, "Post retrieved successfully", post)
}

// GetUserPostsController handles retrieving posts by a specific user
// @Summary Get posts by user ID
// @Description Retrieve paginated posts created by a specific user
// @Tags posts
// @Produce json
// @Param userId path int true "User ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page (max 100)" default(10)
// @Success 200 {object} response.APIResponse{data=dto.PostsListResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/users/{userId}/posts [get]
func (pc *PostController) GetUserPostsController(c echo.Context) error {
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return response.ErrorResponse(c, apperrors.ErrInvalidInput)
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	posts, err := pc.postService.GetUserPosts(uint(userID), page, limit)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, 200, "Posts retrieved successfully", posts)
}

// GetAllPostsController handles retrieving all posts
// @Summary Get all posts
// @Description Retrieve paginated list of all posts
// @Tags posts
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page (max 100)" default(10)
// @Success 200 {object} response.APIResponse{data=dto.PostsListResponse}
// @Failure 500 {object} response.APIResponse
// @Router /api/posts [get]
func (pc *PostController) GetAllPostsController(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	posts, err := pc.postService.GetAllPosts(page, limit)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, 200, "Posts retrieved successfully", posts)
}

// UpdatePostController handles updating a post
// @Summary Update a post
// @Description Update an existing post (only by post owner)
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body dto.UpdatePostRequest true "Updated post data"
// @Success 200 {object} response.APIResponse{data=dto.PostResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Security ApiKeyAuth
// @Router /api/posts/{id} [put]
func (pc *PostController) UpdatePostController(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.ErrorResponse(c, apperrors.ErrInvalidInput)
	}

	var req dto.UpdatePostRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, apperrors.ErrInvalidInput)
	}

	if err := validation.ValidateStruct(&req); err != nil {
		return response.ErrorResponse(c, err)
	}

	// Get user ID from JWT token
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return response.ErrorResponse(c, apperrors.ErrUnauthorized)
	}

	post, err := pc.postService.UpdatePost(uint(id), userID, &req)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, 200, "Post updated successfully", post)
}

// DeletePostController handles deleting a post
// @Summary Delete a post
// @Description Delete an existing post (only by post owner)
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Security ApiKeyAuth
// @Router /api/posts/{id} [delete]
func (pc *PostController) DeletePostController(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.ErrorResponse(c, apperrors.ErrInvalidInput)
	}

	// Get user ID from JWT token
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return response.ErrorResponse(c, apperrors.ErrUnauthorized)
	}

	if err := pc.postService.DeletePost(uint(id), userID); err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, 200, "Post deleted successfully", nil)
}
