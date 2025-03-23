// [The adapter pattern for middleware](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-adapter-pattern-for-middleware)
package server

import (
	"log"
	"net/http"
	"time"
)

//TODO middleware and middleware grouping

// loggingMiddleware is middleware that logs requests
func loggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logger.Printf("%s %s Started: %s", r.Method, r.URL.Path, time.Since(start))
			next.ServeHTTP(w, r)
			logger.Printf("%s %s Ended:%s", r.Method, r.URL.Path, time.Since(start))
		},
	)
}

func corsMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

			// Handle preflight OPTIONS requests
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			logger.Printf("yep did cors")

			// Proceed with the next handler
			next.ServeHTTP(w, r)
		},
	)
}

// note how this chains middlewares together but also handles dependencies so that calling code doesn't have to.
func newMiddleCore(
	logger *log.Logger,
) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		// Apply middlewares in reverse order - last one is applied first
		return loggingMiddleware(logger,
			corsMiddleware(logger, h),
		)
	}
}
