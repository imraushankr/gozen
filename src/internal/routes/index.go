package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imraushankr/gozen/src/internal/config"
	"github.com/imraushankr/gozen/src/pkg/response"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// API v1 group
	apiV1 := app.Group("/api/v1")
	
	// Add health endpoint under API v1 group
	apiV1.Get("/health", func(c *fiber.Ctx) error {
		return response.Success(c, "API is running", fiber.Map{
			"app":     cfg.App.Name,
			"version": cfg.App.Version,
			"env":     cfg.App.Environment,
		})
	})
	
	// Setup route groups
	SetupAuthRoutes(apiV1, cfg)
	SetupUserRoutes(apiV1, cfg)
	// Add more route groups as needed
	// SetupAdminRoutes(apiV1, cfg)
	// SetupProductRoutes(apiV1, cfg)
}