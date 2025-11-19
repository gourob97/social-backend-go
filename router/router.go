package router

import (
    "social-backend/controller"

    "github.com/labstack/echo/v4"
)
type Router struct {
    echo *echo.Echo
    userController *controller.UserController
}

func NewRouter(echo *echo.Echo, userController *controller.UserController) *Router {
    return &Router{
        echo: echo,
        userController: userController,
    }
}

func (r *Router) SetupRoutes() {
    e := r.echo
    api := e.Group("/api")
    api.POST("/register", r.userController.RegisterUser)
    api.POST("/login", r.userController.LoginUser)
}