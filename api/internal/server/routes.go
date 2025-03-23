package server

import (
	"api/internal/database"
	"log"
	"net/http"
)

// TODO https://www.reddit.com/r/golang/comments/1jcmfrb/supermuxer_tiny_and_compact_dependencyfree/

// addRoutes maps all the API routes
// [Map the entire API surface in routes.go](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo)

func addRoutes(mux *http.ServeMux, logger *log.Logger, db *database.Postgres) {

	middleCore := newMiddleCore(logger)

	mux.Handle(http.MethodGet+" /api/clients", middleCore(handleListClients(logger, db)))
	mux.Handle(http.MethodGet+" /api/clients/{id}", middleCore(handleGetClient(logger, db)))
	mux.Handle(http.MethodPost+" /api/clients", middleCore(handleCreateClient(logger, db)))
	mux.Handle(http.MethodPut+" /api/clients/{id}", middleCore(handleUpdateClient(logger, db)))
	// TODO PATCH
	mux.Handle(http.MethodDelete+" /api/clients/{id}", middleCore(handleDeleteClient(logger, db)))
	mux.Handle(http.MethodGet+" /healthz", middleCore(handleHealthz()))
	mux.Handle(http.MethodGet+" /healthdbz", middleCore(handleHealthDBz(logger, db)))
	mux.Handle("/", middleCore(http.NotFoundHandler()))
}
