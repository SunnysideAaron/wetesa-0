package server

import (
	"errors"
	"log/slog"
	"net/http"
	"runtime/debug"

	"api/internal/logging"
)

// handleErrorExample demonstrates what the code looks like for handling errors
func handleErrorExample(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.LogAttrs(
				r.Context(),
				slog.LevelDebug,
				"test log entry",
			)

			err := errors.New("i'm a demon sent to torment you 😈")

			stack := logging.FormatStack(debug.Stack())

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
