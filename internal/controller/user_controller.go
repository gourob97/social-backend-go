package controller

import (
    "social-backend/internal/service"

    "github.com/labstack/echo/v4"

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

func RegisterController(c echo.Context) error {
    var req RegisterRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(400, map[string]string{"error": "cannot parse body"})
    }

    if err := service.RegisterUser(req.Username, req.Email, req.Password); err != nil {
        return c.JSON(400, map[string]string{"error": err.Error()})
    }

    return c.JSON(201, map[string]string{"message": "User registered successfully!"})
}

func LoginController(c echo.Context) error {
    var req LoginRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(400, map[string]string{"error": "cannot parse body"})
    }

    user, err := service.LoginUser(req.Email, req.Password)
    if err != nil {
        return c.JSON(401, map[string]string{"error": "Invalid email or password"})
    }
    return c.JSON(200, map[string]interface{}{
        "message": "Login successful",
        "user":    user,
    })
}