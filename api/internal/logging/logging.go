// TODO error stack traces
// https://betterstack.com/community/guides/logging/logging-in-go/#error-logging-with-slog
// provided example includes a dependency which breaks the goal of this project.
// punting for now.
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
		Level:     slog.LevelDebug,
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
	l     *log.Logger
	attrs []slog.Attr // Add this field to store attributes
}

// slogFields is the key for storing slog attributes in context
type ctxKey string

const slogFields = ctxKey("slog-fields")

// AppendCtx adds slog attributes to context
// https://betterstack.com/community/guides/logging/logging-in-go/#using-the-context-package-with-slog
func AppendCtx(ctx context.Context, attrs ...slog.Attr) context.Context {
	if existing, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		attrs = append(existing, attrs...)
	}
	return context.WithValue(ctx, slogFields, attrs)
}

func getValue(v slog.Value) interface{} {
	if v.Kind() == slog.KindAny {
		if logValuer, ok := v.Any().(interface{ LogValue() slog.Value }); ok {
			return getValue(logValuer.LogValue())
		}
	}
	return v.Any()
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	// Add any context attributes to the record
	// https://betterstack.com/community/guides/logging/logging-in-go/#using-the-context-package-with-slog
	if attrs, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

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

	fields := make(map[string]interface{})

	// Add the handler's stored attrs
	for _, a := range h.attrs {
		fields[a.Key] = getValue(a.Value)
	}

	// Add the record's attrs
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = getValue(a.Value)
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
		attrs:   make([]slog.Attr, 0), // Initialize attrs slice
	}
	return h
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Create a new handler with combined attributes
	newAttrs := make([]slog.Attr, len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)
	copy(newAttrs[len(h.attrs):], attrs)

	return &PrettyHandler{
		Handler: h.Handler.WithAttrs(attrs),
		l:       h.l,
		attrs:   newAttrs,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// Create a new handler with the same logger but with an additional group
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
