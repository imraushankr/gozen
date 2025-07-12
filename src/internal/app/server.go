package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gozen/src/configs"
	"gozen/src/internal/pkg/database"
	"gozen/src/internal/pkg/logger"

	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
	db         *database.DB
	cfg        *configs.Config
}

func NewServer(cfg *configs.Config) *Server {
	// Initialize database
	db, err := database.NewDB(&cfg.Database)
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	// Initialize router with dependencies
	router := SetupRouter(cfg, db)

	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
			Handler:      router,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
		},
		db:  db,
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	logger.Info("Starting server",
		zap.String("host", s.cfg.Server.Host),
		zap.String("port", s.cfg.Server.Port),
		zap.String("environment", s.cfg.App.Environment),
	)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	if err := s.db.Close(); err != nil {
		logger.Error("Failed to close database", zap.Error(err))
	}

	logger.Info("Server shutdown complete")
	return nil
}

func (s *Server) Run() {
	// Start server
	if err := s.Start(); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
		return
	}

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	if err := s.Stop(); err != nil {
		logger.Fatal("Failed to shutdown server", zap.Error(err))
	}
}
