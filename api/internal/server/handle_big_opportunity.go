package server

import (
	"log/slog"
	"net/http"
	"time"
)

// handleBigOpportunity demonstrates what the code in routes will look like if
// things need to take a long time to process.
// There are no problems just opportunities... yeah right.
func handleBigOpportunity(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(35 * time.Second)

			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("OK I Finished"))
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"could not write OK response",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
}
