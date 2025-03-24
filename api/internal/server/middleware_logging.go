package server

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the start and end of a request
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
