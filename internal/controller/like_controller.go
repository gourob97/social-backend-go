package controller

import (
	"strconv"

	"social-backend/internal/response"
	"social-backend/internal/service/interfaces"

	"github.com/labstack/echo/v4"
)

type LikeController struct {
	likeService interfaces.LikeService
}

func NewLikeController(likeService interfaces.LikeService) *LikeController {
	return &LikeController{
		likeService: likeService,
	}
}

// ToggleLikeController handles POST /api/posts/:id/like
// @Summary Toggle like on a post
// @Description Like or unlike a post (toggle like status)
// @Tags likes
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} response.APIResponse{data=dto.LikeResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Security ApiKeyAuth
// @Router /api/posts/{id}/like [post]
func (lc *LikeController) ToggleLikeController(c echo.Context) error {
	// Get post ID from URL parameter
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid post ID")
	}

	// Get user ID from JWT context (extracted from token by middleware)
	userID := c.Get("userID").(uint)

	// Toggle like
	likeResponse, err := lc.likeService.ToggleLike(userID, uint(postID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.OK(c, "Like status updated successfully", likeResponse)
}

// GetLikeStatusController handles GET /api/posts/:id/like
// @Summary Get like status for a post
// @Description Get the like count and whether the current user has liked the post
// @Tags likes
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} response.APIResponse{data=dto.LikeResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/posts/{id}/like [get]
func (lc *LikeController) GetLikeStatusController(c echo.Context) error {
	// Get post ID from URL parameter
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		return response.BadRequest(c, "Invalid post ID")
	}

	// Get user ID from JWT context (0 if not authenticated)
	userID := uint(0)
	if userIDVal := c.Get("userID"); userIDVal != nil {
		userID = userIDVal.(uint)
	}

	// Get like status
	likeResponse, err := lc.likeService.GetPostLikeStatus(userID, uint(postID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.OK(c, "Like status retrieved successfully", likeResponse)
}
