package main

import (
	"github.com/imraushankr/gozen/src/internal/config"
	"github.com/imraushankr/gozen/src/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Convert config.LoggerConfig to logger.Config
	loggerConfig := &logger.Config{
		Level:      cfg.Logger.Level,
		Format:     cfg.Logger.Format,
		Output:     cfg.Logger.Output,
		FilePath:   cfg.Logger.FilePath,
		MaxSize:    cfg.Logger.MaxSize,
		MaxBackups: cfg.Logger.MaxBackups,
		MaxAge:     cfg.Logger.MaxAge,
		Compress:   cfg.Logger.Compress,
	}

	// Initialize logger with config
	if err := logger.InitLogger(loggerConfig); err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()

	// Log application startup
	logger.Info("Starting application",
		zap.String("name", cfg.App.Name),
		zap.String("version", cfg.App.Version),
		zap.String("environment", cfg.App.Environment),
		zap.String("server", cfg.Server.Host+":"+cfg.Server.Port),
	)

	// Example of using configuration
	logger.Info("Database configuration",
		zap.String("type", cfg.Database.Type),
		zap.String("database", cfg.Database.Database),
		zap.Int("max_open_conns", cfg.Database.MaxOpenConns),
	)

	// Example of using email configuration
	if cfg.Email.Provider != "" {
		logger.Info("Email configuration loaded",
			zap.String("provider", cfg.Email.Provider),
			zap.String("smtp_host", cfg.Email.SMTP.Host),
			zap.Int("smtp_port", cfg.Email.SMTP.Port),
			zap.Bool("use_tls", cfg.Email.SMTP.UseTLS),
		)
	}

	// Your application logic here...
	
	logger.Info("Application started successfully")
}