package handler

import (
    "social-backend/internal/service"

    "github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func RegisterHandler(c *fiber.Ctx) error {
    var req RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse body"})
    }

    if err := service.RegisterUser(req.Username, req.Email, req.Password); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully!"})
}

func LoginHandler(c *fiber.Ctx) error {
    var req LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse body"})
    }

    user, err := service.LoginUser(req.Email, req.Password)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
    }

    return c.JSON(fiber.Map{"message": "Welcome " + user.Username})
}