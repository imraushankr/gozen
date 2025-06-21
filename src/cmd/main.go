package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/imraushankr/gozen/src/app"
	"github.com/imraushankr/gozen/src/internal/config"
	"github.com/imraushankr/gozen/src/internal/db"
	"github.com/imraushankr/gozen/src/internal/initialize"
	"github.com/imraushankr/gozen/src/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Setup the application
	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize dependencies (this should initialize the logger)
	if err := initialize.Initialize(cfg); err != nil {
		log.Fatal("Failed to initialize dependencies:", err)
	}

	// Setup Fiber app
	fiberApp := app.SetupApp(cfg)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		logger.Info("Shutting down server...")

		// Shutdown with timeout
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
		defer cancel()

		if err := fiberApp.ShutdownWithContext(ctx); err != nil {
			logger.Error("Server forced to shutdown", zap.Error(err))
		}

		// Close database connections
		if err := db.Close(); err != nil {
			logger.Error("Failed to close database connections", zap.Error(err))
		}

		// Sync logger
		logger.Sync()
	}()

	// Start server
	addr := cfg.GetServerAddress()
	logger.Info("Starting server", zap.String("address", addr))

	if err := fiberApp.Listen(addr); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}