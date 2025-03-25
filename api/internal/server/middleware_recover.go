package server

import (
	"log/slog"
	"net/http"
)

// Copied from https://github.com/google/exposure-notifications-server/blob/main/internal/middleware/recovery.go
// This is simple and should cover us for now.
// If what ever I choose for logging / error handling later doesn't give a stack
// trace or enough info try one of the others.
// https://github.com/labstack/echo/blob/master/middleware/recover.go
// https://github.com/go-chi/chi/blob/v1.5.5/middleware/recoverer.go#L21
func recoverMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					logger.Error("panic recovered",
						"request_id", requestIDFromContext(r.Context()),
						"method", r.Method,
						"path", r.URL.Path,
						"panic", p,
					)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		},
	)
}
