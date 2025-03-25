package server

import (
	"log/slog"
	"net/http"
	"time"
)

// LoggingMiddleware logs the start and end of a request
func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			requestID := requestIDFromContext(r.Context())

			logger.Info("request started",
				"request_id", requestID,
				"method", r.Method,
				"path", r.URL.Path,
			)

			next.ServeHTTP(w, r)

			logger.Info("request completed",
				"request_id", requestID,
				"method", r.Method,
				"path", r.URL.Path,
				"duration", time.Since(start),
			)
		},
	)
}
