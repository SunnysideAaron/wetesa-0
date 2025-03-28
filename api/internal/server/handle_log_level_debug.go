package server

import (
	"log/slog"
	"net/http"
)

func handleLogLevelDebug(logger *slog.Logger, logLevel *slog.LevelVar) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logLevel.Set(slog.LevelDebug)

			logger.LogAttrs(
				r.Context(),
				slog.LevelDebug,
				"set log level to debug",
			)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		},
	)
}
