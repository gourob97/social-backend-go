package router

import (
	"social-backend/internal/controller"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userController *controller.UserController) {
	api := e.Group("/api")
	api.POST("/register", userController.RegisterController)
	api.POST("/login", userController.LoginController)
}
