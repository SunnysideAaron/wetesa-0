package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"api/internal/logging"
)

// RequestIDKey is the context key used to store and retrieve the request ID.
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
func generateID() (string, error) {
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		_, err := rand.Read(buf[:])
		if err != nil {
			return "", fmt.Errorf("could not generate id: %w", err)
		}

		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}
	return b64[0:10], nil
}

// requestID is an HTTP middleware that ensures each request has a unique ID.
// If a request ID is not already present in the context, it generates a new one using
// generateID() and adds it to the context. This ID can be used for request tracing
// and correlation across logs and services.
//
// The middleware will:
// 1. Check for an existing request ID in the context
// 2. Generate a new ID if none exists
// 3. Add the ID to the request context
// 4. Return a 500 error if ID generation fails
func requestID(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if existing := requestIDFromContext(ctx); existing == "" {
				id, err := generateID()

				if err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelInfo,
						"could not generate request id",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				ctx = logging.AppendCtx(ctx, slog.String(RequestIDKey, id))
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		},
	)
}
