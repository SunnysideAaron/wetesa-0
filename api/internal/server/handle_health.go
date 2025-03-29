package server

import (
	"api/internal/database"
	"log/slog"
	"net/http"
)

// handleHealthz handles api server health check requests
func handleHealthz() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		},
	)
}

// handleHealthDBz handles database health check requests
func handleHealthDBz(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			status := http.StatusServiceUnavailable
			stats := db.Health(r.Context(), logger)
			if stats["status"] == "up" {
				status = http.StatusOK
			}

			err := encode(w, r, status, stats)
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
