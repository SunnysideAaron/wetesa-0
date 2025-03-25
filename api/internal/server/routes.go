package server

import (
	"api/internal/config"
	"api/internal/database"
	"log/slog"
	"net/http"
)

// AddRoutes maps all the API routes
// [Map the entire API surface in routes.go](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo)
func AddRoutes(logger *slog.Logger, cfg *config.APIConfig, db *database.Postgres) http.Handler {
	baseMux := http.NewServeMux()
	v1Mux := http.NewServeMux()

	middleDefaults := newMiddleDefaults(logger, cfg)
	middleCore := newMiddleCore(logger)

	v1Mux.Handle(http.MethodGet+" /clients", middleDefaults(handleListClients(logger, db)))
	v1Mux.Handle(http.MethodGet+" /clients/{id}", middleDefaults(handleGetClient(logger, db)))
	v1Mux.Handle(http.MethodPost+" /clients", middleDefaults(middleCore(handleCreateClient(logger, db))))
	v1Mux.Handle(http.MethodPut+" /clients/{id}", middleDefaults(handleUpdateClient(logger, db)))
	v1Mux.Handle(http.MethodDelete+" /clients/{id}", middleDefaults(handleDeleteClient(logger, db)))

	// TODO how to do breaking changes to an api. WARNING hot wire topic but something has to be done.
	baseMux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1Mux))
	baseMux.Handle(http.MethodGet+" /healthz", middleDefaults(handleHealthz()))
	baseMux.Handle(http.MethodGet+" /healthdbz", middleDefaults(handleHealthDBz(logger, db)))
	baseMux.Handle("/", middleDefaults(http.NotFoundHandler()))

	return baseMux
}
