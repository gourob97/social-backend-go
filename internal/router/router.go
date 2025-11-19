package router

import (
	"time"
	"social-backend/internal/controller"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userController *controller.UserController) {
	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status":    "OK",
			"timestamp": time.Now(),
			"service":   "Social Backend Go",
			"version":   "v1.0.0",
		})
	})
	
	api := e.Group("/api")
	api.POST("/register", userController.RegisterController)
	api.POST("/login", userController.LoginController)
}
