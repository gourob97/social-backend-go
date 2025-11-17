package router

import (
    "social-backend/internal/controller"

    "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
    api := e.Group("/api")
    api.POST("/register", controller.RegisterController)
    api.POST("/login", controller.LoginController)
}