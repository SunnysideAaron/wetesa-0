package server

import (
	"api/internal/logging"
	"context"
	"log/slog"
	"net/http"
	"strings"
)

const IPKey string = "ip_address"

// IPFromContext pulls the IP address from the context, if one was set.
// If one was not set, it returns the empty string.
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

func ipMiddleware(next http.Handler) http.Handler {
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
