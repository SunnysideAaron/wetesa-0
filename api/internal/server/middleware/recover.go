package middleware

import (
	"log/slog"
	"net/http"
)

// panicRecovery recovers from panics in HTTP handlers, logs the error,
// and returns a 500 Internal Server Error response to the client.
// Copied from https://github.com/google/exposure-notifications-server/blob/main/internal/middleware/recovery.go
// This is simple and should cover us for now.
// If what ever I choose for logging / error handling later doesn't give a stack
// trace or enough info try one of the others.
// https://github.com/labstack/echo/blob/master/middleware/recover.go
// https://github.com/go-chi/chi/blob/v1.5.5/middleware/recoverer.go#L21
func panicRecovery(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelError,
						"panic recovered",
						slog.String("request_id", requestIDFromContext(r.Context())),
						slog.String("method", r.Method),
						slog.String("path", r.URL.Path),
						slog.Any(
							"panic",
							p,
						), // TODO verify if this slog.Any will actually log. pretty handler might ignore slog.Any
					)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		},
	)
}
