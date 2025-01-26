package lumina

import (
	"context"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

type ContextKey string

const (
	RequestIDKey ContextKey = "requestID"
	UserIDKey    ContextKey = "userID"
)

type LuminaIface interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)

	DebugF(format string, args ...interface{})
	InfoF(format string, args ...interface{})
	WarnF(format string, args ...interface{})
	ErrorF(format string, args ...interface{})

	DebugFCtx(ctx context.Context, format string, args ...interface{})
	InfoFCtx(ctx context.Context, format string, args ...interface{})
	WarnFCtx(ctx context.Context, format string, args ...interface{})
	ErrorFCtx(ctx context.Context, format string, args ...interface{})
}

type luminaLogger struct {
	Log *zap.Logger
}

// NewLogger creates and returns a new Logger instance that implements the LuminaIface interface.
// Returns:
//   - LuminaIface: A Logger instance that implements the LuminaIface interface.
//   - error: An error if the logger creation fails, or nil if successful.
//
// Example:
//
//	logger, err := NewLogger()
//	if err != nil {
//	    log.Fatalf("Failed to create logger: %v", err)
//	}
//	logger.Info("Logger initialized successfully") // Use the logger
func NewLogger() (LuminaIface, error) {
	logger, err := createLogger()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
