package server

import (
	"errors"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"
)

// formatStack makes the stack trace more readable by:
// - Removing unnecessary runtime info
// - Removing extra blank lines
func formatStack(stack []byte) string {
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

// handleErrorExample demonstrates what the code looks like for handling errors
func handleErrorExample(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := errors.New("i'm a demon sent to torment you")
			//err = fmt.Errorf("error occurred: %w", err)

			// Create a separate field for stack trace
			stack := formatStack(debug.Stack())

			logger.LogAttrs(
				r.Context(),
				slog.LevelError,
				err.Error(),
				slog.String("stack_trace", stack),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	)
}
