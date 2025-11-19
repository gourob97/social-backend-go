package middleware

import (
	"strings"

	"social-backend/internal/apperrors"
	"social-backend/internal/auth"
	"social-backend/internal/response"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	jwtManager *auth.JWTManager
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager: auth.NewJWTManager(jwtSecret),
	}
}

func (am *AuthMiddleware) JWTAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return response.ErrorResponse(c, apperrors.ErrUnauthorized)
			}

			// Check if the header starts with "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return response.ErrorResponse(c, apperrors.ErrUnauthorized)
			}

			// Extract the token
			tokenString := authHeader[7:] // Remove "Bearer " prefix

			// Validate the token
			claims, err := am.jwtManager.ValidateToken(tokenString)
			if err != nil {
				return response.ErrorResponse(c, apperrors.ErrUnauthorized)
			}

			// Set user information in context
			c.Set("userID", claims.UserID)
			c.Set("username", claims.Username)
			c.Set("email", claims.Email)

			return next(c)
		}
	}
}
