package database

import (
	"database/sql"
	"gorm.io/gorm"
)

// DB holds both SQL and GORM database connections
type DB struct {
	SQL  *sql.DB
	GORM *gorm.DB
}