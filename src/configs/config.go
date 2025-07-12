package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Email    EmailConfig    `mapstructure:"email"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	CORS     CORSConfig     `mapstructure:"cors"`
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
	Debug       bool   `mapstructure:"debug"`
}

type ServerConfig struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type DatabaseConfig struct {
	Type            string        `mapstructure:"type"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Database        string        `mapstructure:"database"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

type JWTConfig struct {
	AccessTokenSecret  string        `mapstructure:"access_token_secret"`
	AccessTokenExpiry  time.Duration `mapstructure:"access_token_expiry"`
	RefreshTokenSecret string        `mapstructure:"refresh_token_secret"`
	RefreshTokenExpiry time.Duration `mapstructure:"refresh_token_expiry"`
	Issuer             string        `mapstructure:"issuer"`
}

type EmailConfig struct {
	Provider string     `mapstructure:"provider"`
	SMTP     SMTPConfig `mapstructure:"smtp"`
}

type SMTPConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	FromEmail string `mapstructure:"from_email"`
	FromName  string `mapstructure:"from_name"`
	UseTLS    bool   `mapstructure:"use_tls"`
}

type LoggerConfig struct {
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	Output     string `mapstructure:"output"`
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type CORSConfig struct {
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age"`
}

func LoadConfig() (*Config, error) {
	// Load .env file (ignore error if file doesn't exist)
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, using system environment variables")
	}

	// Set up Viper
	viper.AutomaticEnv()
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath(filepath.Join("src", "configs"))

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Process all configuration values with substitutions
	processSubstitutions()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func processSubstitutions() {
	// Helper function to get substituted value
	getValue := func(key string) string {
		val := viper.GetString(key)
		if strings.HasPrefix(val, "${") && strings.HasSuffix(val, "}") {
			envVar := val[2 : len(val)-1]
			envVal := os.Getenv(envVar)
			if envVal != "" {
				return envVal
			}
			return ""
		}
		return val
	}

	// Process all values that need substitution
	viper.Set("app.name", getValue("app.name"))
	viper.Set("app.environment", getValue("app.environment"))
	viper.Set("app.debug", parseBool(getValue("app.debug"), false))

	viper.Set("server.host", getValue("server.host"))
	viper.Set("server.port", getValue("server.port"))
	viper.Set("server.read_timeout", parseDuration(getValue("server.read_timeout"), 10*time.Second))
	viper.Set("server.write_timeout", parseDuration(getValue("server.write_timeout"), 10*time.Second))
	viper.Set("server.shutdown_timeout", parseDuration(getValue("server.shutdown_timeout"), 15*time.Second))

	// Database section
	viper.Set("database.type", getValue("database.type"))
	viper.Set("database.host", getValue("database.host"))
	viper.Set("database.port", parseInt(getValue("database.port"), 5432))
	viper.Set("database.username", getValue("database.username"))
	viper.Set("database.password", getValue("database.password"))
	viper.Set("database.database", getValue("database.database"))
	viper.Set("database.ssl_mode", getValue("database.ssl_mode"))
	viper.Set("database.max_open_conns", parseInt(getValue("database.max_open_conns"), 25))
	viper.Set("database.max_idle_conns", parseInt(getValue("database.max_idle_conns"), 5))
	viper.Set("database.conn_max_lifetime", parseDuration(getValue("database.conn_max_lifetime"), 5*time.Minute))

	// JWT section
	viper.Set("jwt.access_token_secret", getValue("jwt.access_token_secret"))
	viper.Set("jwt.access_token_expiry", parseDuration(getValue("jwt.access_token_expiry"), 15*time.Minute))
	viper.Set("jwt.refresh_token_secret", getValue("jwt.refresh_token_secret"))
	viper.Set("jwt.refresh_token_expiry", parseDuration(getValue("jwt.refresh_token_expiry"), 168*time.Hour))
	viper.Set("jwt.issuer", getValue("jwt.issuer"))

	// Email section
	viper.Set("email.provider", getValue("email.provider"))
	viper.Set("email.smtp.host", getValue("email.smtp.host"))
	viper.Set("email.smtp.port", parseInt(getValue("email.smtp.port"), 587))
	viper.Set("email.smtp.username", getValue("email.smtp.username"))
	viper.Set("email.smtp.password", getValue("email.smtp.password"))
	viper.Set("email.smtp.from_email", getValue("email.smtp.from_email"))
	viper.Set("email.smtp.from_name", getValue("email.smtp.from_name"))
	viper.Set("email.smtp.use_tls", parseBool(getValue("email.smtp.use_tls"), true))

	// Logger section
	viper.Set("logger.level", getValue("logger.level"))
	viper.Set("logger.format", getValue("logger.format"))
	viper.Set("logger.output", getValue("logger.output"))
	viper.Set("logger.file_path", getValue("logger.file_path"))
	viper.Set("logger.max_size", parseInt(getValue("logger.max_size"), 100))
	viper.Set("logger.max_backups", parseInt(getValue("logger.max_backups"), 3))
	viper.Set("logger.max_age", parseInt(getValue("logger.max_age"), 28))
	viper.Set("logger.compress", parseBool(getValue("logger.compress"), true))

	// CORS section
	viper.Set("cors.allow_origins", parseStringSlice(getValue("cors.allow_origins")))
	viper.Set("cors.allow_methods", parseStringSlice(getValue("cors.allow_methods")))
	viper.Set("cors.allow_headers", parseStringSlice(getValue("cors.allow_headers")))
	viper.Set("cors.allow_credentials", parseBool(getValue("cors.allow_credentials"), true))
	viper.Set("cors.max_age", parseInt(getValue("cors.max_age"), 300))
}

func parseBool(val string, defaultValue bool) bool {
	if val == "" {
		return defaultValue
	}
	result, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}
	return result
}

func parseInt(val string, defaultValue int) int {
	if val == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return result
}

func parseDuration(val string, defaultValue time.Duration) time.Duration {
	if val == "" {
		return defaultValue
	}
	result, err := time.ParseDuration(val)
	if err != nil {
		return defaultValue
	}
	return result
}

func parseStringSlice(val string) []string {
	if val == "" {
		return nil
	}
	return strings.Split(val, ",")
}