package server

import (
	"log/slog"
	"net/http"
	"strings"
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
			requestID := requestIDFromContext(r.Context())

			logger.LogAttrs(
				r.Context(),
				slog.LevelDebug,
				"logging middleware entered",
				slog.String("path", r.URL.Path),
				slog.String("method", r.Method),
			)

			// Do not log the request body. It may contain sensitive information.
			logger.LogAttrs(
				r.Context(),
				slog.LevelInfo,
				"request started",
				slog.String("ip", getIP(r)),
				slog.String("request_id", requestID),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("referer", r.Header.Get("Referer")),
				slog.String("userAgent", r.Header.Get("User-Agent")),
			)

			wrapped := &responseWriter{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(wrapped, r)

			// Do not log the response body. It may contain sensitive information.
			logger.LogAttrs(
				r.Context(),
				slog.LevelInfo,
				"request completed",
				slog.String("ip", getIP(r)),
				slog.String("request_id", requestID),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", wrapped.status),
				slog.String("duration", time.Since(start).String()),
				slog.Int64("size", wrapped.size),
			)
		},
	)
}

// getIP returns the client's IP address from the request,
// checking X-Forwarded-For and X-Real-IP headers first
func getIP(r *http.Request) string {
	// Check X-Forwarded-For header
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Get the first IP in case of multiple forwards
		return strings.Split(forwarded, ",")[0]
	}

	// Check X-Real-IP header
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}

	// Fall back to RemoteAddr
	return strings.Split(r.RemoteAddr, ":")[0]
}
