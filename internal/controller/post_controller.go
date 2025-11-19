package controller

import (
	"net/http"
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

	return response.SuccessResponse(c, http.StatusCreated, "Post created successfully", post)
}

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

	return response.SuccessResponse(c, http.StatusOK, "Post retrieved successfully", post)
}

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

	return response.SuccessResponse(c, http.StatusOK, "Posts retrieved successfully", posts)
}

func (pc *PostController) GetAllPostsController(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	posts, err := pc.postService.GetAllPosts(page, limit)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Posts retrieved successfully", posts)
}

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

	return response.SuccessResponse(c, http.StatusOK, "Post updated successfully", post)
}

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

	return response.SuccessResponse(c, http.StatusOK, "Post deleted successfully", nil)
}
