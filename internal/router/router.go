package router

import (
    "social-backend/internal/handler"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")
    api.Post("/register", handler.RegisterHandler)
    api.Post("/login", handler.LoginHandler)
}