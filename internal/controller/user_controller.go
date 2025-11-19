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
