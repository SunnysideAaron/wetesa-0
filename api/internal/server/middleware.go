// Package server does the api work.
// middleware includes groups of middleware.
//
// [The adapter pattern for middleware]
// (https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-adapter-pattern-for-middleware)
package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"api/internal/config"
)

// newMiddleCore returns a middleware chain that applies core middleware functions
// in the correct order. It handles cross-cutting concerns like logging, recovery,
// and CORS.
// note how this chains middlewares together but also handles dependencies so
// that calling code doesn't have to.
func newMiddleCore(
	logger *slog.Logger,
) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		// Apply middlewares in reverse order - last one is applied first
		return requestIDMiddleware(logger,
			ipMiddleware(
				loggingMiddleware(logger,
					recoverMiddleware(logger,
						corsMiddleware(
							http.AllowQuerySemicolons(h),
						),
					),
				),
			),
		)
	}
}

// newMiddleDefaults returns a middleware chain that applies default middleware
// functions with configurable timeouts and request size limits.
func newMiddleDefaults(
	ctx context.Context,
	cfg *config.APIConfig,
	logger *slog.Logger,
	opts ...int,
) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		timeout := cfg.APIDefaultWriteTimeout
		maxBytes := cfg.RequestMaxBytes

		// Override defaults if parameters are provided
		if len(opts) > 0 && opts[0] > 0 {
			if opts[0] > int(cfg.APIWriteTimeout) {
				logger.LogAttrs(
					ctx,
					slog.LevelWarn,
					"passed in timeout is greater than the max timeout, using max timeout",
					slog.Int("timeout", opts[0]),
					slog.Int("max_timeout", int(cfg.APIWriteTimeout)),
				)

				opts[0] = int(cfg.APIWriteTimeout)
			}
			timeout = time.Duration(opts[0])
		}
		if len(opts) > 1 && opts[1] > 0 {
			maxBytes = int64(opts[1])
		}

		return http.TimeoutHandler(
			http.MaxBytesHandler(h, maxBytes),
			timeout,
			"request took too long",
		)
	}
}
