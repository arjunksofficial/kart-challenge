package logger

// use zap logger for structured logging
import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *CustomLogger = nil // global logger instance
)

type CustomLogger struct {
	*zap.Logger
}

func InitLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	var err error
	zapLogger, err := config.Build()
	if err != nil {
		panic(err)
	}
	logger = &CustomLogger{zapLogger}
}
func GetLogger() *CustomLogger {
	if logger == nil {
		InitLogger()
	}
	return logger
}

// Close cleans up the logger resources
func Close() {
	if logger != nil {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}
}

// Info logs an info message
func Info(msg string, fields ...zap.Field) {
	if logger == nil {
		InitLogger()
	}
	logger.Info(msg, fields...)
}

func Infof(msg string, args ...interface{}) {
	if logger == nil {
		InitLogger()
	}
	logger.Info(fmt.Sprintf(msg, args...))
}

// Error logs an error message
func Error(msg string, fields ...zap.Field) {
	if logger == nil {
		InitLogger()
	}
	logger.Error(msg, fields...)
}

// Errorf logs an error message with formatted string
func Errorf(msg string, args ...interface{}) {
	if logger == nil {
		InitLogger()
	}
	logger.Error(fmt.Sprintf(msg, args...))
}

// Debug logs a debug message
func Debug(msg string, fields ...zap.Field) {
	if logger == nil {
		InitLogger()
	}
	logger.Debug(msg, fields...)
}

// Warn logs a warning message
func Warn(msg string, fields ...zap.Field) {
	if logger == nil {
		InitLogger()
	}
	logger.Warn(msg, fields...)
}

// Fatal logs a fatal message and exits the application
func Fatal(msg string, fields ...zap.Field) {
	if logger == nil {
		InitLogger()
	}
	logger.Fatal(msg, fields...)
}

// Panic logs a panic message and panics
func Panic(msg string, fields ...zap.Field) {
	if logger == nil {
		InitLogger()
	}
	logger.Panic(msg, fields...)
}

// WithFields logs a message with additional fields
func WithFields(msg string, fields map[string]interface{}) {
	if logger == nil {
		InitLogger()
	}
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	logger.Info(msg, zapFields...)
}

func (logger *CustomLogger) Info(msg string, fields ...zap.Field) {
	logger.Logger.Info(msg, fields...)
}

func (logger *CustomLogger) Infof(msg string, args ...interface{}) {
	logger.Logger.Info(fmt.Sprintf(msg, args...))
}

func (logger *CustomLogger) Error(msg string, fields ...zap.Field) {
	logger.Logger.Error(msg, fields...)
}

func (logger *CustomLogger) Errorf(msg string, args ...interface{}) {
	logger.Logger.Error(fmt.Sprintf(msg, args...))
}

func (logger *CustomLogger) Debug(msg string, fields ...zap.Field) {
	logger.Logger.Debug(msg, fields...)
}

func (logger *CustomLogger) Warn(msg string, fields ...zap.Field) {
	logger.Logger.Warn(msg, fields...)
}

func (logger *CustomLogger) Fatal(msg string, fields ...zap.Field) {
	logger.Logger.Fatal(msg, fields...)
}

func (logger *CustomLogger) Panic(msg string, fields ...zap.Field) {
	logger.Logger.Panic(msg, fields...)
}

// WithFields logs a message with additional fields
func (logger *CustomLogger) WithFields(msg string, fields map[string]interface{}) {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	logger.Logger.Info(msg, zapFields...)
}
func StringField(key string, value string) zap.Field {
	return zap.String(key, value)
}
func IntField(key string, value int) zap.Field {
	return zap.Int(key, value)
}
func BoolField(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}
func Float64Field(key string, value float64) zap.Field {
	return zap.Float64(key, value)
}
func ErrorField(key string, err error) zap.Field {
	return zap.Error(err)
}
func AnyField(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}
