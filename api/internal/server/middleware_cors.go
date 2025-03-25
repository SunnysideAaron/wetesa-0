package server

import (
	"log/slog"
	"net/http"
)

func corsMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
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
