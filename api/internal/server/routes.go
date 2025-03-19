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
	mux.Handle("GET    /api/clients", handleListClients(logger, db))
	mux.Handle("GET    /api/clients/{id}", handleGetClient(logger, db))
	mux.Handle("POST   /api/clients", handleCreateClient(logger, db))
	mux.Handle("PUT    /api/clients/{id}", handleUpdateClient(logger, db))
	mux.Handle("DELETE /api/clients/{id}", handleDeleteClient(logger, db))
	mux.Handle("GET    /healthz", handleHealthz())
	mux.Handle("GET    /healthdbz", handleHealthDBz(logger, db))
	mux.Handle("/", http.NotFoundHandler())
}
