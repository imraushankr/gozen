// package config

// import (
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"gopkg.in/yaml.v3"
// )

// type Config struct {
// 	Server   ServerConfig   `yaml:"server"`
// 	Database DatabaseConfig `yaml:"database"`
// 	JWT      JWTConfig      `yaml:"jwt"`
// }

// type ServerConfig struct {
// 	Host string `yaml:"host"`
// 	Port string `yaml:"port"`
// }

// type DatabaseConfig struct {
// 	Type     string `yaml:"type"` // sqlite, mysql, postgresql, mongodb
// 	Host     string `yaml:"host"`
// 	Port     string `yaml:"port"`
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// 	Database string `yaml:"database"`
// 	SSLMode  string `yaml:"ssl_mode"`
// }

// type JWTConfig struct {
// 	AccessTokenSecret  string        `yaml:"access_token_secret"`
// 	AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`
// 	RefreshTokenSecret string        `yaml:"refresh_token_secret"`
// 	RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"`
// }

// func LoadConfig() *Config {
// 	// Load .env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("No .env file found, using system environment variables")
// 	}

// 	config := &Config{}

// 	// Try to load from config.yaml first
// 	if data, err := os.ReadFile("config.yaml"); err != nil {
// 		if err := yaml.Unmarshal(data, config); err != nil {
// 			log.Fatal("Error parsing config.yaml: ", err)
// 		}
// 	}

// 	// Override with environment variable if present
// 	if host := os.Getenv("SERVER_HOST"); host != "" {
// 		config.Server.Host = host
// 	} else if config.Server.Host == "" {
// 		config.Server.Host = "localhost"
// 	}

// 	if port := os.Getenv("SERVER_PORT"); port != "" {
// 		config.Server.Port = port
// 	} else if config.Server.Port == "" {
// 		config.Server.Port = "3000"
// 	}

// 	if dbType := os.Getenv("DB_TYPE"); dbType != "" {
// 		config.Database.Type = dbType
// 	} else if config.Database.Type == "" {
// 		config.Database.Type = "sqlite"
// 	}

// 	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
// 		config.Database.Host = dbHost
// 	}

// 	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
// 		config.Database.Port = dbPort
// 	}

// 	if dbUser := os.Getenv("DB_USERNAME"); dbUser != "" {
// 		config.Database.Username = dbUser
// 	}

// 	if dbPass := os.Getenv("DB_PASSWORD"); dbPass != "" {
// 		config.Database.Password = dbPass
// 	}

// 	if dbName := os.Getenv("DB_DATABASE"); dbName != "" {
// 		config.Database.Database = dbName
// 	} else if config.Database.Database == "" {
// 		config.Database.Database = "gozen"
// 	}

// 	if sslMode := os.Getenv("DB_SSL_MODE"); sslMode != "" {
// 		config.Database.SSLMode = sslMode
// 	}

// 	// JWT Configuration
// 	if secret := os.Getenv("ACCESS_TOKEN_SECRET"); secret != "" {
// 		config.JWT.AccessTokenSecret = secret
// 	} else if config.JWT.AccessTokenSecret == "" {
// 		config.JWT.AccessTokenSecret = "your-secret-access-token-key"
// 	}

// 	if expiry := os.Getenv("ACCESS_TOKEN_EXPIRY"); expiry != "" {
// 		if duration, err := time.ParseDuration(expiry); err == nil {
// 			config.JWT.AccessTokenExpiry = duration
// 		}
// 	} else if config.JWT.AccessTokenExpiry == 0 {
// 		config.JWT.AccessTokenExpiry = 15 * time.Minute
// 	}

// 	if secret := os.Getenv("REFRESH_TOKEN_SECRET"); secret != "" {
// 		config.JWT.RefreshTokenSecret = secret
// 	} else if config.JWT.RefreshTokenSecret == "" {
// 		config.JWT.RefreshTokenSecret = "your-secret-refresh-token-key"
// 	}

// 	if expiry := os.Getenv("REFRESH_TOKEN_EXPIRY"); expiry != "" {
// 		if duration, err := time.ParseDuration(expiry); err == nil {
// 			config.JWT.RefreshTokenExpiry = duration
// 		}
// 	} else if config.JWT.RefreshTokenExpiry == 0 {
// 		config.JWT.RefreshTokenExpiry = 7 * 24 * time.Hour
// 	}
// 	return config
// }


// config/config.go
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
}

type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"` // development, staging, production
	Debug       bool   `yaml:"debug"`
}

type ServerConfig struct {
	Host            string        `yaml:"host"`
	Port            string        `yaml:"port"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type DatabaseConfig struct {
	Type            string        `yaml:"type"` // sqlite, mysql, postgresql, mongodb
	Host            string        `yaml:"host"`
	Port            string        `yaml:"port"`
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
	Level      string `yaml:"level"`       // debug, info, warn, error
	Format     string `yaml:"format"`     // json, console
	Output     string `yaml:"output"`     // stdout, stderr, file
	FilePath   string `yaml:"file_path"`  // log file path when output is file
	MaxSize    int    `yaml:"max_size"`   // max size in MB
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`    // days
	Compress   bool   `yaml:"compress"`
}

var globalConfig *Config

func LoadConfig(configPath ...string) (*Config, error) {
	path := "config/config.yaml"
	if len(configPath) > 0 && configPath[0] != "" {
		path = configPath[0]
	}

	// Check if config file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	// Set defaults
	setDefaults(config)

	// Validate configuration
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	globalConfig = config
	return config, nil
}

func GetConfig() *Config {
	if globalConfig == nil {
		panic("configuration not loaded. Call LoadConfig() first")
	}
	return globalConfig
}

func setDefaults(config *Config) {
	// App defaults
	if config.App.Name == "" {
		config.App.Name = "GoZen"
	}
	if config.App.Version == "" {
		config.App.Version = "1.0.0"
	}
	if config.App.Environment == "" {
		config.App.Environment = "development"
	}

	// Server defaults
	if config.Server.Host == "" {
		config.Server.Host = "localhost"
	}
	if config.Server.Port == "" {
		config.Server.Port = "3000"
	}
	if config.Server.ReadTimeout == 0 {
		config.Server.ReadTimeout = 10 * time.Second
	}
	if config.Server.WriteTimeout == 0 {
		config.Server.WriteTimeout = 10 * time.Second
	}
	if config.Server.ShutdownTimeout == 0 {
		config.Server.ShutdownTimeout = 15 * time.Second
	}

	// Database defaults
	if config.Database.Type == "" {
		config.Database.Type = "sqlite"
	}
	if config.Database.Database == "" {
		config.Database.Database = "gozen.db"
	}
	if config.Database.MaxOpenConns == 0 {
		config.Database.MaxOpenConns = 25
	}
	if config.Database.MaxIdleConns == 0 {
		config.Database.MaxIdleConns = 5
	}
	if config.Database.ConnMaxLifetime == 0 {
		config.Database.ConnMaxLifetime = 5 * time.Minute
	}

	// JWT defaults
	if config.JWT.AccessTokenExpiry == 0 {
		config.JWT.AccessTokenExpiry = 15 * time.Minute
	}
	if config.JWT.RefreshTokenExpiry == 0 {
		config.JWT.RefreshTokenExpiry = 7 * 24 * time.Hour
	}
	if config.JWT.Issuer == "" {
		config.JWT.Issuer = config.App.Name
	}

	// Logger defaults
	if config.Logger.Level == "" {
		if config.App.Environment == "development" {
			config.Logger.Level = "debug"
		} else {
			config.Logger.Level = "info"
		}
	}
	if config.Logger.Format == "" {
		if config.App.Environment == "development" {
			config.Logger.Format = "console"
		} else {
			config.Logger.Format = "json"
		}
	}
	if config.Logger.Output == "" {
		config.Logger.Output = "stdout"
	}
	if config.Logger.MaxSize == 0 {
		config.Logger.MaxSize = 100
	}
	if config.Logger.MaxBackups == 0 {
		config.Logger.MaxBackups = 3
	}
	if config.Logger.MaxAge == 0 {
		config.Logger.MaxAge = 28
	}
}

func validateConfig(config *Config) error {
	// Validate required JWT secrets
	if config.JWT.AccessTokenSecret == "" {
		return fmt.Errorf("JWT access token secret is required")
	}
	if config.JWT.RefreshTokenSecret == "" {
		return fmt.Errorf("JWT refresh token secret is required")
	}

	// Validate email config if provider is set
	if config.Email.Provider != "" {
		if config.Email.SMTP.Host == "" {
			return fmt.Errorf("SMTP host is required when email provider is set")
		}
		if config.Email.SMTP.Username == "" {
			return fmt.Errorf("SMTP username is required when email provider is set")
		}
		if config.Email.SMTP.Password == "" {
			return fmt.Errorf("SMTP password is required when email provider is set")
		}
		if config.Email.SMTP.FromEmail == "" {
			return fmt.Errorf("SMTP from_email is required when email provider is set")
		}
	}

	return nil
}