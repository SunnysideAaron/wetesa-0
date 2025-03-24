// [The adapter pattern for middleware](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-adapter-pattern-for-middleware)
package server

import (
	"api/internal/config"
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
	"time"
)

const RequestIDKey string = "request_id"

// requestIDFromContext pulls the request ID from the context, if one was set.
// If one was not set, it returns the empty string.
func requestIDFromContext(ctx context.Context) string {
	v := ctx.Value(RequestIDKey)
	if v == nil {
		return ""
	}

	t, ok := v.(string)
	if !ok {
		return ""
	}
	return t
}

// generateID creates a random 10-character string using base64 encoding
// this is copied from https://github.com/go-chi/chi/blob/v1.5.5/middleware/request_id.go#L67
//
// A quick note on the statistics here: we're trying to calculate the chance that
// two randomly generated base62 prefixes will collide. We use the formula from
// http://en.wikipedia.org/wiki/Birthday_problem
//
// P[m, n] \approx 1 - e^{-m^2/2n}
//
// We ballpark an upper bound for $m$ by imagining (for whatever reason) a server
// that restarts every second over 10 years, for $m = 86400 * 365 * 10 = 315360000$
//
// For a $k$ character base-62 identifier, we have $n(k) = 62^k$
//
// Plugging this in, we find $P[m, n(10)] \approx 5.75%$, which is good enough for
// our purposes, and is surely more than anyone would ever need in practice -- a
// process that is rebooted a handful of times a day for a hundred years has less
// than a millionth of a percent chance of generating two colliding IDs.
func generateID() string {
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		rand.Read(buf[:])
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}
	return b64[0:10]
}

func requestIDMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if existing := requestIDFromContext(ctx); existing == "" {
				id := generateID()
				ctx = context.WithValue(ctx, RequestIDKey, id)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		},
	)
}

// loggingMiddleware logs the start and end of a request
func loggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logger.Printf("%s %s %s Started: %s", requestIDFromContext(r.Context()), r.Method, r.URL.Path, time.Since(start))
			next.ServeHTTP(w, r)
			logger.Printf("%s %s %s Ended:%s", requestIDFromContext(r.Context()), r.Method, r.URL.Path, time.Since(start))
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
		return requestIDMiddleware(logger,
			loggingMiddleware(logger,
				corsMiddleware(logger,
					http.AllowQuerySemicolons(h),
				),
			),
		)
	}
}

func newMiddleDefaults(
	logger *log.Logger,
	cfg *config.APIConfig,
) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		core := newMiddleCore(logger)

		return core(
			http.TimeoutHandler(
				http.MaxBytesHandler(h, cfg.RequestMaxBytes),
				cfg.RequestTimeout*time.Second,
				"request took too long",
			),
		)
	}
}
