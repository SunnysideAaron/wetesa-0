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
			if err := encode(w, r, http.StatusOK, db.Health(r.Context())); err != nil {
				logger.Error("error encoding response", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		},
	)
}
