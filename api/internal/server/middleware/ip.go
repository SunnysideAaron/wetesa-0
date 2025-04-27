package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"api/internal/logging"
)

// IPKey is the context key used to store and retrieve the client's IP address.
// This key is used both in the context and in structured logging.
const IPKey string = "ip_address"

// ipFromContext pulls the IP address from the context, if one was set.
// If one was not set, it returns the empty string.
//
// This function is used to avoid duplicate IP address lookups and ensure
// consistent IP reporting throughout the request lifecycle.
func ipFromContext(ctx context.Context) string {
	v := ctx.Value(IPKey)
	if v == nil {
		return ""
	}

	t, ok := v.(string)
	if !ok {
		return ""
	}
	return t
}

// getIP returns the client's IP address from the request by checking multiple headers
// in order of preference:
//  1. X-Forwarded-For header (first IP if multiple are present)
//  2. X-Real-IP header
//  3. Request's RemoteAddr
//
// This function handles common proxy and load balancer scenarios where the original
// client IP is forwarded in headers. For X-Forwarded-For, only the first IP is used
// as it represents the original client, while subsequent IPs are proxy addresses.
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

// ip is an HTTP middleware that ensures each request has the client's IP
// address available in its context. The IP address is also added to the request's
// structured logging context.
//
// The middleware:
// 1. Checks if an IP is already in the context
// 2. If not, determines the client IP using getIP()
// 3. Adds the IP to both the request context and logging context
//
// This middleware should be placed early in the middleware chain to ensure the IP
// is available for subsequent middleware and request handlers.
func ip(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if existing := ipFromContext(ctx); existing == "" {
				ip := getIP(r)
				ctx = logging.AppendCtx(ctx, slog.String(IPKey, ip))
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		},
	)
}
