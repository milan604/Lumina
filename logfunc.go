package lumina

import (
	"context"
	"fmt"
)

func (l *luminaLogger) Debug(args ...interface{}) {
	l.Log.Debug(args...)
}

func (l *luminaLogger) Info(args ...interface{}) {
	l.Log.Info(args...)
}

func (l *luminaLogger) Warn(args ...interface{}) {
	l.Log.Warn(args...)
}

func (l *luminaLogger) Error(args ...interface{}) {
	l.Log.Error(args...)
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
	keyPairs := withContext(ctx)
	l.Log.With(keyPairs...).Debug(fmt.Sprintf(format, args...))
}

func (l *luminaLogger) InfoFCtx(ctx context.Context, format string, args ...interface{}) {
	keyPairs := withContext(ctx)
	l.Log.With(keyPairs...).Info(fmt.Sprintf(format, args...))
}

func (l *luminaLogger) WarnFCtx(ctx context.Context, format string, args ...interface{}) {
	keyPairs := withContext(ctx)
	l.Log.With(keyPairs...).Warn(fmt.Sprintf(format, args...))
}

func (l *luminaLogger) ErrorFCtx(ctx context.Context, format string, args ...interface{}) {
	keyPairs := withContext(ctx)
	l.Log.With(keyPairs...).Error(fmt.Sprintf(format, args...))
}

func (l *luminaLogger) Sync() error {
	return l.Log.Sync()
}

func (l *luminaLogger) With(fields ...any) LuminaIface {
	return &luminaLogger{
		Log: l.Log.With(fields...),
	}
}
