package server

import (
	"log/slog"
	"net/http"
)

func handleLogLevel(logger *slog.Logger, logLevel *slog.LevelVar) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// level := r.URL.Query().Get("level")
			level := r.PathValue("level")
			var newLevel slog.Level

			switch level {
			case "debug":
				newLevel = slog.LevelDebug
			default:
				newLevel = slog.LevelInfo
			}

			logLevel.Set(newLevel)

			logger.LogAttrs(
				r.Context(),
				newLevel,
				"set log level",
				slog.String("level", newLevel.String()),
			)

			w.WriteHeader(http.StatusOK)

			_, err := w.Write([]byte("OK"))
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
