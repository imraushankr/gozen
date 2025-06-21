package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/imraushankr/gozen/src/pkg/response"
)

// RateLimitMiddleware creates a rate limiting middleware
func RateLimitMiddleware(max int, expiration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return response.Error(c, fiber.StatusTooManyRequests, "Too many requests", nil)
		},
	})
}

// AuthRateLimitMiddleware applies stricter rate limiting for auth endpoints
func AuthRateLimitMiddleware() fiber.Handler {
	return RateLimitMiddleware(5, 15*time.Minute) // 5 requests per 15 minutes
}

// APIRateLimitMiddleware applies general rate limiting for API endpoints
func APIRateLimitMiddleware() fiber.Handler {
	return RateLimitMiddleware(100, 1*time.Minute) // 100 requests per minute
}
