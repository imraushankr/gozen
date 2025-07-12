package database

import (
	"database/sql"

	"github.com/imraushankr/gozen/src/configs"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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