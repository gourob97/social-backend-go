package main

import (
	"os"
    "log"
    "social-backend/internal/config"
    "social-backend/internal/router"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    config.ConnectDB()

    app := fiber.New()
    app.Use(logger.New())

    router.SetupRoutes(app)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default
    }

    log.Fatal(app.Listen(":" + port))
}