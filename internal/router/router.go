package router

import (
	"social-backend/internal/controller"
	"social-backend/internal/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userController *controller.UserController, postController *controller.PostController, authMiddleware *middleware.AuthMiddleware) {
	api := e.Group("/api")

	// Public routes
	api.POST("/register", userController.RegisterController)
	api.POST("/login", userController.LoginController)

	// Public post routes (read-only)
	api.GET("/posts", postController.GetAllPostsController)
	api.GET("/posts/:id", postController.GetPostController)
	api.GET("/users/:userId/posts", postController.GetUserPostsController)

	// Protected routes (require authentication)
	protected := api.Group("", authMiddleware.JWTAuthMiddleware())
	protected.POST("/posts", postController.CreatePostController)
	protected.PUT("/posts/:id", postController.UpdatePostController)
	protected.DELETE("/posts/:id", postController.DeletePostController)
}
