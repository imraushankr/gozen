package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/imraushankr/gozen/src/internal/config"
	"github.com/imraushankr/gozen/src/internal/db"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	setupRoutes(app, db, cfg)

	// Start server
	log.Printf("Server starting on %s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(app.Listen(cfg.Server.Host + ":" + cfg.Server.Port))
}

func setupRoutes(app *fiber.App, db *db.Database, cfg *config.Config) {
	api := app.Group("/api/v1")

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":   "ok",
			"database": db.Type,
		})
	})

	// Auth routes would go here
	// auth := api.Group("/auth")
	// auth.Post("/register", registerHandler)
	// auth.Post("/login", loginHandler)
	// auth.Post("/refresh", refreshTokenHandler)
	// auth.Post("/logout", logoutHandler)
	// auth.Post("/forgot-password", forgotPasswordHandler)
	// auth.Post("/reset-password", resetPasswordHandler)
}