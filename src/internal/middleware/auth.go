package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/imraushankr/gozen/src/internal/config"
	"github.com/imraushankr/gozen/src/internal/models"
	"github.com/imraushankr/gozen/src/internal/security"
	"github.com/imraushankr/gozen/src/pkg/response"
)

// AuthMiddleware validates JWT tokens
func AuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Unauthorized(c, "Authorization header is required")
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer " {
			return response.Unauthorized(c, "Invalid authorization header format")
		}

		token := tokenParts[1]
		claims, err := security.ValidateAccessToken(token, cfg)
		if err != nil {
			return response.Unauthorized(c, "Invalid or expired token")
		}

		// Store user information in context
		c.Locals("user_id", claims.ID)
		c.Locals("user_role", claims.Role)
		c.Locals("user_email", claims.Email)

		return c.Next()
	}
}

// AdminMiddleware ensures the user has admin role
func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("user_role").(models.Role)
		if !ok {
			return response.Forbidden(c, "Access denied")
		}

		if userRole != models.ADMIN {
			return response.Forbidden(c, "Admin access required")
		}

		return c.Next()
	}
}

// OptionalAuthMiddleware validates JWT if present but doesn't require it
func OptionalAuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Next()
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Next()
		}

		token := tokenParts[1]
		claims, err := security.ValidateAccessToken(token, cfg)
		if err != nil {
			return c.Next()
		}

		// Store user information in context
		c.Locals("user_id", claims.ID)
		c.Locals("user_role", claims.Role)
		c.Locals("user_email", claims.Email)

		return c.Next()
	}
}