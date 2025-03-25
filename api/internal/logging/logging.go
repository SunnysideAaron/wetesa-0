package logging

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"

	"api/internal/config"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// colorize returns the string wrapped in ANSI color codes
func colorize(s string, color string) string {
	return color + s + colorReset
}

// Level represents logging levels
type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelError Level = "ERROR"
)

// NewLogger creates a new structured logger.
// Uses PrettyHandler for development and JSONHandler for production.
func NewLogger(cfg *config.APIConfig) *slog.Logger {
	opts := slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}

	var handler slog.Handler
	switch cfg.Environment {
	case config.EnvironmentDev:
		prettyOpts := PrettyHandlerOptions{
			SlogOpts: opts,
		}
		handler = NewPrettyHandler(os.Stdout, prettyOpts)
	default:
		// Default to production logging if environment is not set
		handler = slog.NewJSONHandler(os.Stdout, &opts)
	}

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

// PrettyHandler came from
// https://betterstack.com/community/guides/logging/logging-in-go/#customizing-slog-handlers
type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = colorize(level, colorPurple)
	case slog.LevelInfo:
		level = colorize(level, colorBlue)
	case slog.LevelWarn:
		level = colorize(level, colorYellow)
	case slog.LevelError:
		level = colorize(level, colorRed)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := colorize(r.Message, colorCyan)

	h.l.Println(timeStr, level, msg, colorize(string(b), colorWhite))

	return nil
}

func NewPrettyHandler(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}
