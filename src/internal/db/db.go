package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/imraushankr/gozen/src/internal/config"
	"github.com/imraushankr/gozen/src/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	SQL   *gorm.DB
	Mongo *mongo.Database
	Type  string
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	db := &Database{Type: cfg.Database.Type}

	switch cfg.Database.Type {
	case "sqlite":
		return setupSQLite(db, cfg)
	case "mysql":
		return setupMySQL(db, cfg)
	case "postgresql":
		return setupPostgreSQL(db, cfg)
	case "mongodb":
		return setupMongoDB(db, cfg)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Database.Type)
	}
}

func setupSQLite(db *Database, cfg *config.Config) (*Database, error) {
	dbPath := cfg.Database.Database
	if dbPath == "" {
		dbPath = "gozen.db"
	}

	sqlDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.SQL = sqlDB

	// Auto migrate
	err = db.SQL.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to SQLite database")
	return db, nil
}

func setupMySQL(db *Database, cfg *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	sqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.SQL = sqlDB

	// Auto migrate
	err = db.SQL.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MySQL database")
	return db, nil
}

func setupPostgreSQL(db *Database, cfg *config.Config) (*Database, error) {
	sslMode := cfg.Database.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Database,
		cfg.Database.Port,
		sslMode,
	)

	sqlDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.SQL = sqlDB

	// Auto migrate
	err = db.SQL.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL database")
	return db, nil
}

func setupMongoDB(db *Database, cfg *config.Config) (*Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
	)

	// If no username/password, use simple connection
	if cfg.Database.Username == "" {
		uri = fmt.Sprintf("mongodb://%s:%s", cfg.Database.Host, cfg.Database.Port)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db.Mongo = client.Database(cfg.Database.Database)

	log.Println("Connected to MongoDB database")
	return db, nil
}