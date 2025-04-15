// Package logging provides structured logging functionality with support for
// different output formats based on the environment.
//
// WARNING: Pretty handler is pretty basic. There are several packages for pretty
// output it's probably better to use one of those. This is just for the example
// without dependencies.
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
	// LevelDebug Only turn on for in-depth troubleshooting.
	LevelDebug Level = "DEBUG"
	// LevelInfo default level in production. Enough information to troubleshoot basic problems.
	LevelInfo Level = "INFO"
	// LevelWarn Create a ticket. Something is wrong and needs fixing. Properly handled errors are info not warn.
	LevelWarn Level = "WARN"
	// LevelError Call someone NOW! Something is wrong and needs immediate fixing.
	LevelError Level = "ERROR"
)

// NewLogger creates a new structured logger configured based on the environment.
// It returns a slog.Logger and a LevelVar for dynamic log level control.
func NewLogger(cfg *config.APIConfig) (*slog.Logger, *slog.LevelVar) {
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelInfo)

	opts := slog.HandlerOptions{
		Level:     lvl,
		AddSource: true,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			// didn't seems to work.
			// if a.Key == "request_id" {
			// 	return slog.String("request_id", "🤖")
			// }
			return a
		},
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
	case LevelInfo:
		return slog.LevelInfo
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

// PrettyHandlerOptions configures the behavior of PrettyHandler.
// Source: https://betterstack.com/community/guides/logging/logging-in-go/#customizing-slog-handlers
type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

// PrettyHandler implements slog.Handler to provide pretty-printed log output
// with colors and formatted JSON attributes.
type PrettyHandler struct {
	slog.Handler
	l     *log.Logger // TODO should this be a log or slog?
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

// getValue extracts the actual value from a slog.Value, handling LogValuer interface
// implementations recursively.
func getValue(v slog.Value) any {
	if v.Kind() == slog.KindAny {
		if logValuer, ok := v.Any().(interface{ LogValue() slog.Value }); ok {
			return getValue(logValuer.LogValue())
		}
	}
	return v.Any()
}

// Handle implements slog.Handler.Handle to process log records.
// It formats the output with colors and structured JSON attributes.
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

	fields := make(map[string]any)

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
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := colorize(r.Message, colorCyan)

	h.l.Println(timeStr, level, msg, colorize(string(b), colorWhite))

	return nil
}

// NewPrettyHandler creates a new PrettyHandler with the specified options.
func NewPrettyHandler(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	levelVar, ok := opts.SlogOpts.Level.(*slog.LevelVar)
	if !ok {
		levelVar = new(slog.LevelVar) // Provide a default if assertion fails
		levelVar.Set(slog.LevelInfo)
	}

	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
		attrs:   make([]slog.Attr, 0),
		level:   levelVar,
	}
	return h
}

// WithAttrs implements slog.Handler.WithAttrs to create a new handler with
// the specified attributes added to the set of attributes that will be
// included with all log records.
func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Create a new handler with combined attributes
	newAttrs := make([]slog.Attr, len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)
	copy(newAttrs[len(h.attrs):], attrs)

	return &PrettyHandler{
		Handler: h.Handler.WithAttrs(attrs),
		l:       h.l,
		attrs:   newAttrs,
		level:   h.level,
	}
}

// WithGroup implements slog.Handler.WithGroup to create a new handler with
// the specified group name added to the set of groups that will be
// included with all log records.
func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// Create a new handler with the same logger but with an additional group
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
		attrs:   h.attrs,
		level:   h.level,
	}
}

// enabled checks if the given log level should be processed
func (h *PrettyHandler) enabled(level slog.Level) bool {
	return level >= h.level.Level()
}

// Level returns the current level of the handler
// func (h *PrettyHandler) Level() slog.Level {
// 	return h.level.Level()
// }

// FormatStack makes the stack trace more readable by:
// - Removing unnecessary runtime info
// - Removing extra blank lines
// - Focusing on application code (paths containing "api/")
//
// The output format is:
// → function_name
//
//	at file/path.go:line_number
func FormatStack(stack []byte) string {
	lines := strings.Split(string(stack), "\n")
	var filtered []string

	for i := range lines {
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
