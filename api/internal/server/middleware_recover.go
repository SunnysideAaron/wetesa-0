package server

import (
	"log"
	"net/http"
)

// Copied from https://github.com/google/exposure-notifications-server/blob/main/internal/middleware/recovery.go
// This is simple and should cover us for now.
// If what ever I choose for logging / error handling later doesn't give a stack
// trace or enough info try one of the others.
// https://github.com/labstack/echo/blob/master/middleware/recover.go
// https://github.com/go-chi/chi/blob/v1.5.5/middleware/recoverer.go#L21
func recoverMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					logger.Printf("%s %s %s Recovered from panic", requestIDFromContext(r.Context()), r.Method, r.URL.Path)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}()

			next.ServeHTTP(w, r)
		},
	)
}
