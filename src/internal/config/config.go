package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Type     string `yaml:"type"` // sqlite, mysql, postgresql, mongodb
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

type JWTConfig struct {
	AccessTokenSecret  string        `yaml:"access_token_secret"`
	AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`
	RefreshTokenSecret string        `yaml:"refresh_token_secret"`
	RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"`
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config := &Config{}

	// Try to load from config.yaml first
	if data, err := os.ReadFile("config.yaml"); err != nil {
		if err := yaml.Unmarshal(data, config); err != nil {
			log.Fatal("Error parsing config.yaml: ", err)
		}
	}

	// Override with environment variable if present
	if host := os.Getenv("SERVER_HOST"); host != "" {
		config.Server.Host = host
	} else if config.Server.Host == "" {
		config.Server.Host = "localhost"
	}

	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	} else if config.Server.Port == "" {
		config.Server.Port = "3000"
	}

	if dbType := os.Getenv("DB_TYPE"); dbType != "" {
		config.Database.Type = dbType
	} else if config.Database.Type == "" {
		config.Database.Type = "sqlite"
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.Database.Host = dbHost
	}

	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.Database.Port = dbPort
	}

	if dbUser := os.Getenv("DB_USERNAME"); dbUser != "" {
		config.Database.Username = dbUser
	}

	if dbPass := os.Getenv("DB_PASSWORD"); dbPass != "" {
		config.Database.Password = dbPass
	}

	if dbName := os.Getenv("DB_DATABASE"); dbName != "" {
		config.Database.Database = dbName
	} else if config.Database.Database == "" {
		config.Database.Database = "gozen"
	}

	if sslMode := os.Getenv("DB_SSL_MODE"); sslMode != "" {
		config.Database.SSLMode = sslMode
	}

	// JWT Configuration
	if secret := os.Getenv("ACCESS_TOKEN_SECRET"); secret != "" {
		config.JWT.AccessTokenSecret = secret
	} else if config.JWT.AccessTokenSecret == "" {
		config.JWT.AccessTokenSecret = "your-secret-access-token-key"
	}

	if expiry := os.Getenv("ACCESS_TOKEN_EXPIRY"); expiry != "" {
		if duration, err := time.ParseDuration(expiry); err == nil {
			config.JWT.AccessTokenExpiry = duration
		}
	} else if config.JWT.AccessTokenExpiry == 0 {
		config.JWT.AccessTokenExpiry = 15 * time.Minute
	}

	if secret := os.Getenv("REFRESH_TOKEN_SECRET"); secret != "" {
		config.JWT.RefreshTokenSecret = secret
	} else if config.JWT.RefreshTokenSecret == "" {
		config.JWT.RefreshTokenSecret = "your-secret-refresh-token-key"
	}

	if expiry := os.Getenv("REFRESH_TOKEN_EXPIRY"); expiry != "" {
		if duration, err := time.ParseDuration(expiry); err == nil {
			config.JWT.RefreshTokenExpiry = duration
		}
	} else if config.JWT.RefreshTokenExpiry == 0 {
		config.JWT.RefreshTokenExpiry = 7 * 24 * time.Hour
	}
	return config
}
