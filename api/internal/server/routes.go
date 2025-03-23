package server

import (
	"api/internal/database"
	"log"
	"net/http"
)

// AddRoutes maps all the API routes
// [Map the entire API surface in routes.go](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo)
func AddRoutes(logger *log.Logger, db *database.Postgres) http.Handler {

	baseMux := http.NewServeMux()
	v1Mux := http.NewServeMux()

	middleCore := newMiddleCore(logger)

	v1Mux.Handle(http.MethodGet+" /clients", middleCore(handleListClients(logger, db)))
	v1Mux.Handle(http.MethodGet+" /clients/{id}", middleCore(handleGetClient(logger, db)))
	v1Mux.Handle(http.MethodPost+" /clients", middleCore(handleCreateClient(logger, db)))
	v1Mux.Handle(http.MethodPut+" /clients/{id}", middleCore(handleUpdateClient(logger, db)))
	v1Mux.Handle(http.MethodDelete+" /clients/{id}", middleCore(handleDeleteClient(logger, db)))

	// TODO how to do breaking changes to an api. WARNING hot wire topic but something has to be done.
	baseMux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1Mux))
	baseMux.Handle(http.MethodGet+" /healthz", middleCore(handleHealthz()))
	baseMux.Handle(http.MethodGet+" /healthdbz", middleCore(handleHealthDBz(logger, db)))
	baseMux.Handle("/", middleCore(http.NotFoundHandler()))

	return baseMux
}
