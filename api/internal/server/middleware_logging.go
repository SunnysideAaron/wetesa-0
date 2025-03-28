package server

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int64
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += int64(size)
	return size, err
}

// LoggingMiddleware logs the start and end of a request
func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// for now i'm not putting method and path in context. to demonstrate .With()
			reqLogger := logger.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
			)

			// Do not log the request body, may contain sensitive information.
			reqLogger.LogAttrs(
				r.Context(),
				slog.LevelInfo,
				"request started",
				slog.String("referer", r.Header.Get("Referer")),
				slog.String("userAgent", r.Header.Get("User-Agent")),
			)

			wrapped := &responseWriter{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(wrapped, r)

			// Do not log the response body, may contain sensitive information.
			reqLogger.LogAttrs(
				r.Context(),
				slog.LevelInfo,
				"request completed",
				slog.Int("status", wrapped.status),
				slog.String("duration", time.Since(start).String()),
				slog.Int64("size", wrapped.size),
			)
		},
	)
}
