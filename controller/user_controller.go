package controller

import (
	"social-backend/services"
	"social-backend/utils/errutil"
	"social-backend/utils/msgutil"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserController struct {
	service *	services.UserService
}

func ForNowUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (controller *UserController) RegisterUser(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, msgutil.RequestBodyParseErrorMessage())
	}

	if err := controller.service.RegisterUser(req.Username, req.Email, req.Password); err != nil {
		if err == errutil.ErrCreatingUser {
			return c.JSON(500, msgutil.RegistrationFailureMessage())
		}
		return c.JSON(500, msgutil.InternalServerErrorMessage())

	}

	return c.JSON(201, msgutil.RegistrationSuccessMessage())
}

func (controller *UserController) LoginUser(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, msgutil.RequestBodyParseErrorMessage())
	}

	user, err := controller.service.LoginUser(req.Email, req.Password)
	if err != nil {
		switch err {
		case errutil.ErrUserNotFound:
			return c.JSON(404, msgutil.UserNotFoundMessage())
		case errutil.ErrInvalidCredentials:
			return c.JSON(401, msgutil.InvalidCredentialsMessage())
		default:
			return c.JSON(500, msgutil.InternalServerErrorMessage())
		}
	}
	return c.JSON(200, msgutil.LoginSuccessMessage(*user))
}
