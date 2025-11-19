package main

import (
	"log"
	"social-backend/internal/config"
	"social-backend/internal/controller"
	"social-backend/internal/middleware"
	"social-backend/internal/repository"
	"social-backend/internal/router"
	"social-backend/internal/service"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

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

	e := echo.New()
	e.Use(middleware.EnhancedLogger())
	e.Use(echomiddleware.Recover())

	router.SetupRoutes(e, userController)

	port := cfg.Port
	if port == "" {
		port = "8080" // default
	}

	// Show cute startup banner
	middleware.PrintStartupBanner(port)

	log.Fatal(e.Start(":" + port))
}
