package server

import (
	"log/slog"
	"net/http"

	"api/internal/database"
)

// handleHealthz returns an http.Handler that responds to health check requests
// with a 200 OK status and "OK" message.
func handleHealthz(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			_, err := w.Write([]byte("OK"))
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"could not write OK response",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
}

// handleHealthDBz returns an http.Handler that checks database connectivity
// and responds with database health status information.
func handleHealthDBz(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			status := http.StatusServiceUnavailable
			healthStatus := db.Health(r.Context(), logger)
			if healthStatus.Status == "up" {
				status = http.StatusOK
			}

			err := encode(w, r, status, healthStatus)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelError,
					"error encoding response",
					slog.String("error", err.Error()),
				)

				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)
}
