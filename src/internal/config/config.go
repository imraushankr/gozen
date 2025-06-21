package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	Email    EmailConfig    `yaml:"email"`
	Logger   LoggerConfig   `yaml:"logger"`
	CORS     CORSConfig     `yaml:"cors"`
}

type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
	Debug       bool   `yaml:"debug"`
}

type ServerConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type DatabaseConfig struct {
	Type            string        `yaml:"type"`
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Username        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	Database        string        `yaml:"database"`
	SSLMode         string        `yaml:"ssl_mode"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

type JWTConfig struct {
	AccessTokenSecret  string        `yaml:"access_token_secret"`
	AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`
	RefreshTokenSecret string        `yaml:"refresh_token_secret"`
	RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"`
	Issuer             string        `yaml:"issuer"`
}

type EmailConfig struct {
	Provider string     `yaml:"provider"`
	SMTP     SMTPConfig `yaml:"smtp"`
}

type SMTPConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	FromEmail string `yaml:"from_email"`
	FromName  string `yaml:"from_name"`
	UseTLS    bool   `yaml:"use_tls"`
}

type LoggerConfig struct {
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
	Output     string `yaml:"output"`
	FilePath   string `yaml:"file_path"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

type CORSConfig struct {
	AllowOrigins     []string `yaml:"allow_origins"`
	AllowMethods     []string `yaml:"allow_methods"`
	AllowHeaders     []string `yaml:"allow_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
	MaxAge           int      `yaml:"max_age"`
}

// LoadConfig loads configuration from YAML file only
func LoadConfig(configPath string) (*Config, error) {
	// Set default config path if not provided
	if configPath == "" {
		configPath = "config.yaml"
	}

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", configPath)
	}

	// Read YAML file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Validate configuration
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return config, nil
}

func validateConfig(config *Config) error {
	// Validate app environment
	validEnvs := []string{"development", "staging", "production"}
	validEnv := false
	for _, env := range validEnvs {
		if config.App.Environment == env {
			validEnv = true
			break
		}
	}
	if !validEnv {
		return fmt.Errorf("invalid environment: %s. Must be one of: %v", config.App.Environment, validEnvs)
	}

	// Validate database type
	// validDBTypes := []string{"sqlite", "mysql", "postgresql", "mongodb"}
	validDBTypes := []string{"sqlite", "mysql", "postgresql", "mongodb", "postgres"}
	validDBType := false
	for _, dbType := range validDBTypes {
		if config.Database.Type == dbType {
			validDBType = true
			break
		}
	}
	if !validDBType {
		return fmt.Errorf("invalid database type: %s. Must be one of: %v", config.Database.Type, validDBTypes)
	}

	// Validate server port
	if config.Server.Port <= 0 || config.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d. Must be between 1 and 65535", config.Server.Port)
	}

	// Validate JWT secrets are not empty
	if config.JWT.AccessTokenSecret == "" {
		return fmt.Errorf("JWT access token secret cannot be empty")
	}
	if config.JWT.RefreshTokenSecret == "" {
		return fmt.Errorf("JWT refresh token secret cannot be empty")
	}

	// Validate logger level
	validLogLevels := []string{"debug", "info", "warn", "error"}
	validLogLevel := false
	for _, level := range validLogLevels {
		if config.Logger.Level == level {
			validLogLevel = true
			break
		}
	}
	if !validLogLevel {
		return fmt.Errorf("invalid logger level: %s. Must be one of: %v", config.Logger.Level, validLogLevels)
	}

	// Validate logger format
	validLogFormats := []string{"json", "console"}
	validLogFormat := false
	for _, format := range validLogFormats {
		if config.Logger.Format == format {
			validLogFormat = true
			break
		}
	}
	if !validLogFormat {
		return fmt.Errorf("invalid logger format: %s. Must be one of: %v", config.Logger.Format, validLogFormats)
	}

	// Validate logger output
	validLogOutputs := []string{"stdout", "stderr", "file"}
	validLogOutput := false
	for _, output := range validLogOutputs {
		if config.Logger.Output == output {
			validLogOutput = true
			break
		}
	}
	if !validLogOutput {
		return fmt.Errorf("invalid logger output: %s. Must be one of: %v", config.Logger.Output, validLogOutputs)
	}

	return nil
}

// Helper method to get server address
func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

// Helper method to check if running in production
func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}

// Helper method to check if running in development
func (c *Config) IsDevelopment() bool {
	return c.App.Environment == "development"
}

// Helper method to check if running in staging
func (c *Config) IsStaging() bool {
	return c.App.Environment == "staging"
}

// Helper method to get database connection string for SQL databases
func (c *Config) GetDatabaseDSN() string {
	switch c.Database.Type {
	case "postgres", "postgresql":
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			c.Database.Host, c.Database.Port, c.Database.Username,
			c.Database.Password, c.Database.Database, c.Database.SSLMode)
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Database.Username, c.Database.Password,
			c.Database.Host, c.Database.Port, c.Database.Database)
	case "sqlite":
		return c.Database.Database
	default:
		return ""
	}
}