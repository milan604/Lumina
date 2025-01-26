package lumina

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createLogger() (*luminaLogger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.TimeKey = "timestamp"

	zapLogger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &luminaLogger{
		Log: zapLogger.Sugar(),
	}, nil
}

func withContext(ctx context.Context) []interface{} {
	var fields []interface{}
	// Extract values from context
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		fields = append(fields, zap.String(string(RequestIDKey), requestID))
	}
	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		fields = append(fields, zap.String(string(UserIDKey), userID))
	}

	// Add other fields as per requirement

	return fields
}
