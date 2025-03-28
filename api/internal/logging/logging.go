// TODO Someday read through this guide and see if PrettyHandler could be better.
// https://github.com/golang/example/blob/master/slog-handler-guide/README.md#the-%60enabled%60-method
package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"runtime"
	"strings"

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
func NewLogger(cfg *config.APIConfig) (*slog.Logger, *slog.LevelVar) {
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelInfo)

	opts := slog.HandlerOptions{
		Level:     lvl,
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

	return slog.New(handler), lvl
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
	attrs []slog.Attr
	level *slog.LevelVar
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
	// Check if we should handle this level
	if !h.enabled(r.Level) {
		return nil
	}

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

	// Add source information if available
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		fields["source"] = fmt.Sprintf("%s:%d", f.File, f.Line)
	}

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
		attrs:   make([]slog.Attr, 0),
		level:   opts.SlogOpts.Level.(*slog.LevelVar),
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

func (h *PrettyHandler) enabled(level slog.Level) bool {
	return level >= h.level.Level()
}

// Level returns the current level of the handler
// func (h *PrettyHandler) Level() slog.Level {
// 	return h.level.Level()
// }

// formatStack makes the stack trace more readable by:
// - Removing unnecessary runtime info
// - Removing extra blank lines
func FormatStack(stack []byte) string {
	lines := strings.Split(string(stack), "\n")
	var filtered []string

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "goroutine") {
			continue
		}

		// Skip runtime frames and created by messages
		if strings.Contains(line, "/usr/local/go/src/runtime/") ||
			strings.Contains(line, "created by") {
			continue
		}

		// If this is a function name line
		if !strings.HasPrefix(line, "/") {
			// Remove pointer addresses and clean up
			line = strings.Split(line, "(")[0]
			line = strings.TrimSpace(line)

			// Include all application code
			if strings.Contains(line, "api/") {
				filtered = append(filtered, "→ "+line)
			}
		} else if len(filtered) > 0 { // If we have a previous function name, add its location
			// Extract file and line number
			parts := strings.Split(line, " ")
			if len(parts) > 0 {
				fileParts := strings.Split(parts[0], "api/")
				if len(fileParts) > 1 {
					filtered[len(filtered)-1] += "\n   at " + fileParts[1]
				}
			}
		}
	}

	if len(filtered) == 0 {
		return "<no relevant stack frames>"
	}

	return strings.Join(filtered, "\n")
}
