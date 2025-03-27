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

// TODO convert to variadic parameters so that request time and max bytes can be overwritten per request

func newMiddleDefaults(
	cfg *config.APIConfig,
) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.TimeoutHandler(
			http.MaxBytesHandler(h, cfg.RequestMaxBytes),
			cfg.RequestTimeout*time.Second,
			"request took too long",
		)
	}
}
