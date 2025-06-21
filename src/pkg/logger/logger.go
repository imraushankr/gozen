// package logger

// import (
// 	"os"

// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

// var Logger *zap.Logger

// func init() {
// 	InitLogger()
// }

// // InitLogger initializes the global logger
// func InitLogger() {
// 	config := zap.NewProductionConfig()

// 	// Set log level based on environment
// 	if os.Getenv("APP_ENV") == "development" {
// 		config = zap.NewDevelopmentConfig()
// 		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
// 	} else {
// 		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
// 	}

// 	// Customize time format
// 	config.EncoderConfig.TimeKey = "timestamp"
// 	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

// 	var err error
// 	Logger, err = config.Build()
// 	if err != nil {
// 		panic("Failed to initialize logger: " + err.Error())
// 	}
// }

// // Info logs an info message
// func Info(msg string, fields ...zap.Field) {
// 	Logger.Info(msg, fields...)
// }

// // Error logs an error message
// func Error(msg string, fields ...zap.Field) {
// 	Logger.Error(msg, fields...)
// }

// // Debug logs a debug message
// func Debug(msg string, fields ...zap.Field) {
// 	Logger.Debug(msg, fields...)
// }

// // Warn logs a warning message
// func Warn(msg string, fields ...zap.Field) {
// 	Logger.Warn(msg, fields...)
// }

// // Fatal logs a fatal message and exits
// func Fatal(msg string, fields ...zap.Field) {
// 	Logger.Fatal(msg, fields...)
// }

// // With creates a child logger with additional fields
// func With(fields ...zap.Field) *zap.Logger {
// 	return Logger.With(fields...)
// }

// // Sync flushes any buffered log entries
// func Sync() {
// 	Logger.Sync()
// }


// logger/logger.go
package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

type Config struct {
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
	Output     string `yaml:"output"`
	FilePath   string `yaml:"file_path"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

// InitLogger initializes the global logger with custom configuration
func InitLogger(config *Config) error {
	// Parse log level
	level, err := parseLogLevel(config.Level)
	if err != nil {
		return fmt.Errorf("invalid log level: %w", err)
	}

	// Create encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Create encoder
	var encoder zapcore.Encoder
	if strings.ToLower(config.Format) == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Create writer syncer
	var writeSyncer zapcore.WriteSyncer
	switch strings.ToLower(config.Output) {
	case "stdout":
		writeSyncer = zapcore.AddSync(os.Stdout)
	case "stderr":
		writeSyncer = zapcore.AddSync(os.Stderr)
	case "file":
		if config.FilePath == "" {
			return fmt.Errorf("file path is required when output is file")
		}
		writeSyncer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.FilePath,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		})
	default:
		return fmt.Errorf("invalid output type: %s", config.Output)
	}

	// Create core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// Create logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return nil
}

// InitDefaultLogger initializes logger with default settings
func InitDefaultLogger() {
	config := &Config{
		Level:      "info",
		Format:     "console",
		Output:     "stdout",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	if err := InitLogger(config); err != nil {
		panic("Failed to initialize default logger: " + err.Error())
	}
}

func parseLogLevel(level string) (zapcore.Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel, nil
	case "info":
		return zap.InfoLevel, nil
	case "warn", "warning":
		return zap.WarnLevel, nil
	case "error":
		return zap.ErrorLevel, nil
	case "fatal":
		return zap.FatalLevel, nil
	default:
		return zap.InfoLevel, fmt.Errorf("unknown log level: %s", level)
	}
}

// Convenience functions
func Info(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Info(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Error(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Debug(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Warn(msg, fields...)
	}
}

func Fatal(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Fatal(msg, fields...)
	}
}

func With(fields ...zap.Field) *zap.Logger {
	if Logger != nil {
		return Logger.With(fields...)
	}
	return nil
}

func Sync() error {
	if Logger != nil {
		return Logger.Sync()
	}
	return nil
}