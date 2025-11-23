package controller

import (
	"strconv"

	"social-backend/internal/dto"
	"social-backend/internal/response"
	"social-backend/internal/service/interfaces"
	"social-backend/internal/validation"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	commentService interfaces.CommentService
}

func NewCommentController(commentService interfaces.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

// CreateCommentController handles POST /api/posts/:postId/comments
// @Summary Create a new comment
// @Description Create a new comment for a specific post
// @Tags comments
// @Accept json
// @Produce json
// @Param postId path int true "Post ID"
// @Param comment body dto.CreateCommentRequest true "Comment data"
// @Success 201 {object} response.APIResponse{data=dto.CommentResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Security ApiKeyAuth
// @Router /api/posts/{postId}/comments [post]
func (cc *CommentController) CreateCommentController(c echo.Context) error {
	// Get post ID from URL parameter
	postIDStr := c.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid post ID")
	}

	// Get user ID from JWT context
	userID := c.Get("userID").(uint)

	// Parse and validate request
	var req dto.CreateCommentRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if err := validation.ValidateStruct(&req); err != nil {
		return response.ErrorResponse(c, err)
	}

	// Create comment
	comment, err := cc.commentService.CreateComment(userID, uint(postID), &req)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.Created(c, "Comment created successfully", comment)
}

// GetCommentController handles GET /api/comments/:id
// @Summary Get a comment by ID
// @Description Retrieve a specific comment by its ID
// @Tags comments
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} response.APIResponse{data=dto.CommentResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/comments/{id} [get]
func (cc *CommentController) GetCommentController(c echo.Context) error {
	// Get comment ID from URL parameter
	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid comment ID")
	}

	// Get comment
	comment, err := cc.commentService.GetComment(uint(commentID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.OK(c, "Comment retrieved successfully", comment)
}

// GetPostCommentsController handles GET /api/posts/:postId/comments
// @Summary Get comments for a post
// @Description Retrieve paginated comments for a specific post
// @Tags comments
// @Produce json
// @Param postId path int true "Post ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page (max 100)" default(10)
// @Success 200 {object} response.APIResponse{data=dto.CommentsListResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/posts/{postId}/comments [get]
func (cc *CommentController) GetPostCommentsController(c echo.Context) error {
	// Get post ID from URL parameter
	postIDStr := c.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid post ID")
	}

	// Get pagination parameters
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Get comments
	comments, err := cc.commentService.GetCommentsByPostID(uint(postID), page, limit)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.OK(c, "Comments retrieved successfully", comments)
}

// UpdateCommentController handles PUT /api/comments/:id
// @Summary Update a comment
// @Description Update an existing comment (only by comment owner)
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param comment body dto.UpdateCommentRequest true "Updated comment data"
// @Success 200 {object} response.APIResponse{data=dto.CommentResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Security ApiKeyAuth
// @Router /api/comments/{id} [put]
func (cc *CommentController) UpdateCommentController(c echo.Context) error {
	// Get comment ID from URL parameter
	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid comment ID")
	}

	// Get user ID from JWT context
	userID := c.Get("userID").(uint)

	// Parse and validate request
	var req dto.UpdateCommentRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if err := validation.ValidateStruct(&req); err != nil {
		return response.ErrorResponse(c, err)
	}

	// Update comment
	comment, err := cc.commentService.UpdateComment(userID, uint(commentID), &req)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.OK(c, "Comment updated successfully", comment)
}

// DeleteCommentController handles DELETE /api/comments/:id
// @Summary Delete a comment
// @Description Delete an existing comment (only by comment owner)
// @Tags comments
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Security ApiKeyAuth
// @Router /api/comments/{id} [delete]
func (cc *CommentController) DeleteCommentController(c echo.Context) error {
	// Get comment ID from URL parameter
	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid comment ID")
	}

	// Get user ID from JWT context
	userID := c.Get("userID").(uint)

	// Delete comment
	if err := cc.commentService.DeleteComment(userID, uint(commentID)); err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.OK(c, "Comment deleted successfully", nil)
}

// GetUserCommentsController handles GET /api/users/:userId/comments
// @Summary Get comments by user
// @Description Retrieve paginated comments created by a specific user
// @Tags comments
// @Produce json
// @Param userId path int true "User ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page (max 100)" default(10)
// @Success 200 {object} response.APIResponse{data=dto.CommentsListResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/users/{userId}/comments [get]
func (cc *CommentController) GetUserCommentsController(c echo.Context) error {
	// Get user ID from URL parameter
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid user ID")
	}

	// Get pagination parameters
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Get comments
	comments, err := cc.commentService.GetCommentsByUserID(uint(userID), page, limit)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.OK(c, "User comments retrieved successfully", comments)
}
