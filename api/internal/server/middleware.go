// [The adapter pattern for middleware](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-adapter-pattern-for-middleware)
package server

import (
	"api/internal/config"
	"log/slog"
	"net/http"
	"time"
)

// note how this chains middlewares together but also handles dependencies so that calling code doesn't have to.
func newMiddleCore(
	logger *slog.Logger,
) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		// Apply middlewares in reverse order - last one is applied first
		return requestIDMiddleware(logger,
			loggingMiddleware(logger,
				recoverMiddleware(logger,
					corsMiddleware(logger,
						http.AllowQuerySemicolons(h),
					),
				),
			),
		)
	}
}

func newMiddleDefaults(
	cfg *config.APIConfig,
	logger *slog.Logger,
	opts ...int,
) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		timeout := cfg.APIDefaultWriteTimeout * time.Second
		maxBytes := cfg.RequestMaxBytes

		// Override defaults if parameters are provided
		if len(opts) > 0 && opts[0] > 0 {
			if opts[0] > int(cfg.APIWriteTimeout) {
				logger.Warn("passed in timeout is greater than the max timeout, using max timeout",
					"timeout", opts[0],
					"max_timeout", cfg.APIWriteTimeout,
				)

				opts[0] = int(cfg.APIWriteTimeout)
			}
			timeout = time.Duration(opts[0]) * time.Second
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
