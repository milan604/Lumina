package lumina

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

func (l *luminaLogger) Debug(msg string, fields ...zap.Field) {
	l.Log.Debug(msg, fields...)
}

func (l *luminaLogger) Info(msg string, fields ...zap.Field) {
	l.Log.Info(msg, fields...)
}

func (l *luminaLogger) Warn(msg string, fields ...zap.Field) {
	l.Log.Warn(msg, fields...)
}

func (l *luminaLogger) Error(msg string, fields ...zap.Field) {
	l.Log.Error(msg, fields...)
}

func (l *luminaLogger) DebugF(format string, args ...interface{}) {
	l.Log.Debug(fmt.Sprintf(format, args...))
}

func (l *luminaLogger) InfoF(format string, args ...interface{}) {
	l.Log.Info(fmt.Sprintf(format, args...))
}

func (l *luminaLogger) WarnF(format string, args ...interface{}) {
	l.Log.Warn(fmt.Sprintf(format, args...))
}

func (l *luminaLogger) ErrorF(format string, args ...interface{}) {
	l.Log.Error(fmt.Sprintf(format, args...))
}

func (l *luminaLogger) DebugFCtx(ctx context.Context, format string, args ...interface{}) {
	l.Log.Debug(fmt.Sprintf(format, args...), withContext(ctx)...)
}

func (l *luminaLogger) InfoFCtx(ctx context.Context, format string, args ...interface{}) {
	l.Log.Info(fmt.Sprintf(format, args...), withContext(ctx)...)
}

func (l *luminaLogger) WarnFCtx(ctx context.Context, format string, args ...interface{}) {
	l.Log.Warn(fmt.Sprintf(format, args...), withContext(ctx)...)
}

func (l *luminaLogger) ErrorFCtx(ctx context.Context, format string, args ...interface{}) {
	l.Log.Error(fmt.Sprintf(format, args...), withContext(ctx)...)
}

func (l *luminaLogger) Sync() error {
	return l.Log.Sync()
}
