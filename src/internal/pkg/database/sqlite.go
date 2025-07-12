package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gozen/src/configs"
)

func newSQLiteDB(cfg *configs.DatabaseConfig) (*DB, error) {
	sqlDB, err := sql.Open("sqlite3", cfg.Database)
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(sqlite.Open(cfg.Database), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{SQL: sqlDB, GORM: gormDB}, nil
}