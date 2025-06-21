package db

import (
  "context"
  "fmt"
  "github.com/imraushankr/gozen/src/internal/config"
  "github.com/imraushankr/gozen/src/internal/models"
  "go.mongodb.org/mongo-driver/mongo"
  "gorm.io/gorm"
)

var (
  SQLInstance   *gorm.DB
  MongoInstance *mongo.Database
)

// Initialize sets up the database connection based on config
func InitDatabase(cfg config.DatabaseConfig) error {
  switch cfg.Type {
  case "sqlite":
    return initSQLite(cfg)
  case "mysql":
    return initMySQL(cfg)
  case "postgresql", "postgres":
    return initPostgreSQL(cfg)
  case "mongodb":
    return initMongoDB(cfg)
  default:
    return fmt.Errorf("unsupported database type: %s", cfg.Type)
  }
}

// Close closes database connections
func Close() error {
  if SQLInstance != nil {
    sqlDB, err := SQLInstance.DB()
    if err != nil {
      return err
    }
    return sqlDB.Close()
  }
  if MongoInstance != nil {
    return MongoInstance.Client().Disconnect(context.Background())
  }
  return nil
}

// AutoMigrate runs database migrations for SQL databases
func AutoMigrate() error {
  if SQLInstance == nil {
    return fmt.Errorf("SQL database not initialized")
  }
  return SQLInstance.AutoMigrate(
    &models.User{},
  )
}