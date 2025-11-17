package main

import (
    "log"
    "social-backend/internal/config"
    "social-backend/internal/router"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    // Load config and connect to database
    cfg := config.LoadConfig()
    config.ConnectDB()

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    router.SetupRoutes(e)

    port := cfg.Port
    if port == "" {
        port = "8080" // default
    }

    log.Fatal(e.Start(":" + port))
}