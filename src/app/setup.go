package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/imraushankr/gozen/src/internal/config"
	"github.com/imraushankr/gozen/src/internal/middleware"
	"github.com/imraushankr/gozen/src/internal/routes"
	"github.com/imraushankr/gozen/src/pkg/response"
)

// Setup creates and configures the Fiber application
func SetupApp(cfg *config.Config) *fiber.App {
	// Create Fiber app with custom config
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return response.Error(c, code, err.Error(), nil)
		},
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.CORSMiddleware(&cfg.CORS))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return response.Success(c, "API is running", fiber.Map{
			"app":     cfg.App.Name,
			"version": cfg.App.Version,
			"env":     cfg.App.Environment,
		})
	})

	// Setup API routes
	routes.SetupRoutes(app, cfg)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return response.NotFound(c, "Route not found")
	})

	return app
}
