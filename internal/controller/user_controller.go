package controller

import (
	"social-backend/internal/dto"
	"social-backend/internal/response"
	"social-backend/internal/service/interfaces"
	"social-backend/internal/validation"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService interfaces.UserService
	validator   *validation.Validator
}

func NewUserController(userService interfaces.UserService) *UserController {
	return &UserController{
		userService: userService,
		validator:   validation.NewValidator(),
	}
}

// RegisterController handles user registration
// @Summary Register a new user
// @Description Create a new user account with username, email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "User registration data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/register [post]
func (ctrl *UserController) RegisterController(c echo.Context) error {
	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request format")
	}

	if err := ctrl.validator.Validate(req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := ctrl.userService.RegisterUser(req.Username, req.Email, req.Password); err != nil {
		return response.Error(c, err)
	}

	return response.Created(c, "User registered successfully", nil)
}

// LoginController handles user authentication
// @Summary Login user
// @Description Authenticate user with email and password, returns JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "User login credentials"
// @Success 200 {object} response.APIResponse{data=dto.LoginResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/login [post]
func (ctrl *UserController) LoginController(c echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request format")
	}

	if err := ctrl.validator.Validate(req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	loginResp, err := ctrl.userService.LoginUser(req.Email, req.Password)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, "Login successful", loginResp)
}
