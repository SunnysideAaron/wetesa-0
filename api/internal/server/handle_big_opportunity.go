package server

import (
	"net/http"
	"time"
)

// handleBigOpportunity demonstrates what the code in routes will look like if
// things need to take a long time to process.
// There are no problems just opportunities... yeah right.
func handleBigOpportunity() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Wait for 70 seconds
			time.Sleep(20 * time.Second)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		},
	)
}
