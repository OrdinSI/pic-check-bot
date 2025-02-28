package log

import (
	"context"
	"log/slog"
	"os"
)

var defaultLogger *slog.Logger
var logFile *os.File

// NewLogger ...
func NewLogger(dev bool) error {

	var handler slog.Handler
	if dev {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	} else {
		logFile, err := os.OpenFile("b2b_trainer.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		handler = slog.NewJSONHandler(logFile, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	defaultLogger = slog.New(handler)
	slog.SetDefault(defaultLogger)

	return nil
}

// CloseLogger ...
func CloseLogger() error {
	if logFile != nil {
		return logFile.Close()
	}
	return nil
}

type Logger struct {
	*slog.Logger
}

// Named ...
func Named(name string) *Logger {
	return &Logger{
		Logger: defaultLogger.With("name", name),
	}
}

// Info ...
func (l *Logger) Info(msg string, args ...interface{}) {
	l.Logger.Info(msg, args...)
}

// InfoC ...
func (l *Logger) InfoC(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.InfoContext(ctx, msg, args...)
}

// Debug ...
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Logger.Debug(msg, args...)
}

// DebugC ...
func (l *Logger) DebugC(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.DebugContext(ctx, msg, args...)
}

// Warn ...
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.Logger.Warn(msg, args...)
}

// WarnC ...
func (l *Logger) WarnC(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.WarnContext(ctx, msg, args...)
}

// Error ...
func (l *Logger) Error(msg string, args ...interface{}) {
	l.Logger.Error(msg, args...)
}

// ErrorC ...
func (l *Logger) ErrorC(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.ErrorContext(ctx, msg, args...)
}

// Fatal ...
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.Logger.Error(msg, args...)
	os.Exit(1)
}

// FatalC ...
func (l *Logger) FatalC(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}
