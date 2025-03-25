package logging

import (
	"log/slog"
	"os"
)

// Level represents logging levels
type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelError Level = "ERROR"
)

// NewLogger creates a new structured logger
func NewLogger() *slog.Logger {
	// Start with JSON handler for structured logging
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		// Add source code location to log entries
		AddSource: true,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	return slog.New(handler)
}

// ParseLevel converts a string level to slog.Level
func ParseLevel(level string) slog.Level {
	switch Level(level) {
	case LevelDebug:
		return slog.LevelDebug
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
