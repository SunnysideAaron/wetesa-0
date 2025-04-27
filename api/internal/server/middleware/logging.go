package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"api/internal/logging"
)

// responseWriter wraps http.ResponseWriter to capture the response status code
// and size for logging purposes.
type responseWriter struct {
	http.ResponseWriter
	status int
	size   int64
	// body   []byte
}

// WriteHeader captures the status code and calls the underlying ResponseWriter's
// WriteHeader method.
func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

// Write captures the response size and calls the underlying ResponseWriter's
// Write method.
func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += int64(size)
	// rw.body = append(rw.body, b...)
	return size, fmt.Errorf("ResponseWriter broke: %w", err)
}

// requestLogging logs the start and completion of each request including
// timing, status code, and response size.
func requestLogging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ctx := r.Context()

			ctx = logging.AppendCtx(ctx, slog.String("method", r.Method))
			ctx = logging.AppendCtx(ctx, slog.String("path", r.URL.Path))
			r = r.WithContext(ctx)

			// Do not log the request body, may contain sensitive information.
			logger.LogAttrs(
				r.Context(),
				slog.LevelInfo,
				"request started",
				slog.String("referer", r.Header.Get("Referer")),
				slog.String("userAgent", r.Header.Get("User-Agent")),
			)

			wrapped := &responseWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
				size:           0,
			}
			next.ServeHTTP(wrapped, r)

			// TODO Do not log the response body, may contain sensitive information.
			// eventually logging will have to output full response. punting for now.
			logger.LogAttrs(
				r.Context(),
				slog.LevelInfo,
				"request completed",
				slog.Int("status", wrapped.status),
				slog.String("duration", time.Since(start).String()),
				slog.Int64("size", wrapped.size),
				// slog.String("response_body", string(wrapped.body)),
			)
		},
	)
}
