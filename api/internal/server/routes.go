// Package server provides HTTP server functionality including request handling,
// middleware, and routing for the API.
package server

import (
	"context"
	"log/slog"
	"net/http"

	"api/internal/config"
	"api/internal/database"
	"api/internal/logging"
	"api/internal/server/middleware"
)

// AddRoutes maps all the API routes
// [Map the entire API surface in routes.go](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo)
func AddRoutes(
	ctx context.Context,
	cfg *config.APIConfig,
	db *database.Postgres,
	logger *slog.Logger,
	logLevel *slog.LevelVar,
) http.Handler {
	baseMux := http.NewServeMux()
	v1Mux := http.NewServeMux()

	middleDefaults := middleware.NewDefaults(ctx, cfg, logger)

	// example of overriding defaults
	v1Mux.Handle(http.MethodGet+" /bigopportunity", middleware.NewDefaults(ctx, cfg, logger, 50)(handleBigOpportunity(logger)))
	// directly callable example of an error
	v1Mux.Handle(http.MethodGet+" /errorexample", middleDefaults(handleErrorExample(logger)))
	v1Mux.Handle(http.MethodGet+" /loglevel/{level}", middleDefaults(handleLogLevel(logger, logLevel)))

	// Example of some code having a different log level.
	clientLogger, clientLogLevel := logging.NewLogger(cfg)
	slog.SetDefault(clientLogger)

	v1Mux.Handle(http.MethodGet+" /clients/loglevel/{level}", middleDefaults(handleLogLevel(clientLogger, clientLogLevel)))
	v1Mux.Handle(http.MethodGet+" /clients", middleDefaults(handleListClients(clientLogger, db)))
	v1Mux.Handle(http.MethodGet+" /clients/{id}", middleDefaults(handleGetClient(clientLogger, db)))
	v1Mux.Handle(http.MethodPost+" /clients", middleDefaults(handleCreateClient(clientLogger, db)))
	v1Mux.Handle(http.MethodPut+" /clients/{id}", middleDefaults(handleUpdateClient(clientLogger, db)))
	v1Mux.Handle(http.MethodDelete+" /clients/{id}", middleDefaults(handleDeleteClient(clientLogger, db)))

	// TODO how to do breaking changes to an api. WARNING hot wire topic but something has to be done.
	baseMux.Handle(cfg.BaseURL+"/", http.StripPrefix(cfg.BaseURL, v1Mux))
	baseMux.Handle(http.MethodGet+" /healthz", middleDefaults(handleHealthz(logger)))

	baseMux.Handle(http.MethodGet+" /healthdbz", middleDefaults(handleHealthDBz(logger, db)))

	// due to how go works middleware directly on NotFoundHandler is never called.
	// have to wrap the mux instead.
	baseMux.Handle("/", http.NotFoundHandler())

	// Wrap the entire baseMux with core middleware
	return middleware.NewCore(logger)(baseMux)
}
