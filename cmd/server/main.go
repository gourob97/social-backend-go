package main

import (
	"log"
	_ "social-backend/docs" // This line is necessary for swag to find your docs!
	"social-backend/internal/config"
	"social-backend/internal/controller"
	"social-backend/internal/middleware"
	"social-backend/internal/repository"
	"social-backend/internal/router"
	"social-backend/internal/service"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Social Backend API
// @version 1.0
// @description This is a social media backend API with user authentication, posts, likes, and comments
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize dependencies with dependency injection
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cfg.JWTSecret)
	userController := controller.NewUserController(userService)

	// Initialize post dependencies
	postRepo := repository.NewPostRepository(db)

	// Initialize like dependencies
	likeRepo := repository.NewLikeRepository(db)

	// Initialize comment dependencies
	commentRepo := repository.NewCommentRepository(db)

	// Initialize services with all dependencies
	postService := service.NewPostService(postRepo, likeRepo, commentRepo)
	likeService := service.NewLikeService(likeRepo, postRepo)
	commentService := service.NewCommentService(commentRepo, postRepo)

	// Initialize controllers
	postController := controller.NewPostController(postService)
	likeController := controller.NewLikeController(likeService)
	commentController := controller.NewCommentController(commentService)

	// Initialize auth middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	e := echo.New()
	e.Use(middleware.EnhancedLogger())
	e.Use(echomiddleware.Recover())

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	router.SetupRoutes(e, userController, postController, likeController, commentController, authMiddleware)

	port := cfg.Port
	if port == "" {
		port = "8080" // default
	}

	// Show cute startup banner
	middleware.PrintStartupBanner(port)

	log.Fatal(e.Start(":" + port))
}
