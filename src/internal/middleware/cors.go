package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/imraushankr/gozen/src/internal/config"
)

func CORSMiddleware(cfg *config.CORSConfig) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     strings.Join(cfg.AllowOrigins, ","),
		AllowMethods:     strings.Join(cfg.AllowMethods, ","),
		AllowHeaders:     strings.Join(cfg.AllowHeaders, ","),
		AllowCredentials: cfg.AllowCredentials,
		MaxAge:           cfg.MaxAge,
	})
}