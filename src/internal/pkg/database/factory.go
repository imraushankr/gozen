package database

import (
	"context"
	"fmt"

	"time"

	"github.com/imraushankr/gozen/src/configs"
	"github.com/imraushankr/gozen/src/internal/pkg/logger"
	"go.uber.org/zap"
)

func NewDB(cfg *configs.DatabaseConfig) (*DB, error) {
	var db *DB
	var err error

	switch cfg.Type {
	case "postgres":
		db, err = newPostgresDB(cfg)
	case "mysql":
		db, err = newMySQLDB(cfg)
	case "sqlite":
		db, err = newSQLiteDB(cfg)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}

	if err != nil {
		return nil, err
	}

	// Configure connection pool
	db.SQL.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SQL.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SQL.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.SQL.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Database connection established",
		zap.String("type", cfg.Type),
		zap.String("database", cfg.Database),
	)

	return db, nil
}