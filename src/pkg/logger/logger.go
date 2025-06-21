package logger

import (
	"fmt"
	"github.com/imraushankr/gozen/src/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

var Logger *zap.Logger

// InitLogger initializes the global logger with custom configuration
func InitLogger(loggerConfig *config.LoggerConfig) error {
	// Parse log level
	level, err := parseLogLevel(loggerConfig.Level)
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
	if strings.ToLower(loggerConfig.Format) == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Create writer syncer
	var writeSyncer zapcore.WriteSyncer
	switch strings.ToLower(loggerConfig.Output) {
	case "stdout":
		writeSyncer = zapcore.AddSync(os.Stdout)
	case "stderr":
		writeSyncer = zapcore.AddSync(os.Stderr)
	case "file":
		if loggerConfig.FilePath == "" {
			return fmt.Errorf("file path is required when output is file")
		}
		writeSyncer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   loggerConfig.FilePath,
			MaxSize:    loggerConfig.MaxSize,
			MaxBackups: loggerConfig.MaxBackups,
			MaxAge:     loggerConfig.MaxAge,
			Compress:   loggerConfig.Compress,
		})
	default:
		return fmt.Errorf("invalid output type: %s", loggerConfig.Output)
	}

	// Create core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// Create logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return nil
}

// InitDefaultLogger initializes logger with default settings
func InitDefaultLogger() {
	loggerConfig := &config.LoggerConfig{
		Level:      "info",
		Format:     "console",
		Output:     "stdout",
		FilePath:   "",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}
	if err := InitLogger(loggerConfig); err != nil {
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