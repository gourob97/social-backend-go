package router

import (
	"social-backend/internal/controller"
	"social-backend/internal/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userController *controller.UserController, postController *controller.PostController, likeController *controller.LikeController, commentController *controller.CommentController, authMiddleware *middleware.AuthMiddleware) {
	api := e.Group("/api")

	// Public routes
	api.POST("/register", userController.RegisterController)
	api.POST("/login", userController.LoginController)

	// Public post routes (read-only)
	api.GET("/posts", postController.GetAllPostsController)
	api.GET("/posts/:id", postController.GetPostController)
	api.GET("/users/:userId/posts", postController.GetUserPostsController)
	api.GET("/posts/:id/like", likeController.GetLikeStatusController)
	api.GET("/posts/:postId/comments", commentController.GetPostCommentsController)
	api.GET("/comments/:id", commentController.GetCommentController)
	api.GET("/users/:userId/comments", commentController.GetUserCommentsController)

	// Protected routes (require authentication)
	protected := api.Group("", authMiddleware.JWTAuthMiddleware())
	protected.POST("/posts", postController.CreatePostController)
	protected.PUT("/posts/:id", postController.UpdatePostController)
	protected.DELETE("/posts/:id", postController.DeletePostController)
	protected.POST("/posts/:id/like", likeController.ToggleLikeController)
	protected.POST("/posts/:postId/comments", commentController.CreateCommentController)
	protected.PUT("/comments/:id", commentController.UpdateCommentController)
	protected.DELETE("/comments/:id", commentController.DeleteCommentController)
}
