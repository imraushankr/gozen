package main

import (
	"flag"
	"os"
	"fmt"
	"log"
	"path/filepath"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go run cmd/migrate/run/main.go [up|down|status|version]")
		os.Exit(1)
	}

	// Get project root and paths
	projectRoot, err := getProjectRoot()
	if err != nil {
		log.Fatalf("Error finding project root: %v", err)
	}

	migrationsPath := filepath.Join(projectRoot, "src", "migrations")
	dbPath := filepath.Join(projectRoot, "gozen.db")

	// Check if migrations directory exists
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		log.Fatalf("Migrations directory not found: %s", migrationsPath)
	}

	fmt.Printf("Using migrations from: %s\n", migrationsPath)
	fmt.Printf("Using database: %s\n", dbPath)

	// Add custom database URL with parameters
	dbURL := fmt.Sprintf("sqlite3://%s?_foreign_keys=on", filepath.ToSlash(dbPath))
	
	m, err := migrate.New(
		fmt.Sprintf("file://%s", filepath.ToSlash(migrationsPath)),
		dbURL,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}
	defer m.Close()

	switch args[0] {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to run migrations up: %v", err)
		}
		fmt.Println("Migrations applied successfully")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to run migrations down: %v", err)
		}
		fmt.Println("Migrations rolled back successfully")
	case "status":
		version, dirty, err := m.Version()
		if err != nil {
			if err == migrate.ErrNilVersion {
				fmt.Println("No migrations applied")
				return
			}
			log.Fatalf("Failed to get migration status: %v", err)
		}
		fmt.Printf("Current version: %d (dirty: %v)\n", version, dirty)
	case "version":
		version, _, err := m.Version()
		if err != nil {
			if err == migrate.ErrNilVersion {
				fmt.Println("No migrations applied")
				return
			}
			log.Fatalf("Failed to get migration version: %v", err)
		}
		fmt.Printf("Current migration version: %d\n", version)
	default:
		fmt.Println("Invalid command. Use 'up', 'down', 'status' or 'version'")
		os.Exit(1)
	}
}

func getProjectRoot() (string, error) {
	// Start from current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}

	// Walk up the directory tree to find go.mod
	current := wd
	for {
		if _, err := os.Stat(filepath.Join(current, "go.mod")); err == nil {
			// Found project root
			return current, nil
		}
		parent := filepath.Dir(current)
		if parent == current {
			// Reached filesystem root
			break
		}
		current = parent
	}

	return "", fmt.Errorf("could not locate project root (go.mod not found in any parent directory)")
}