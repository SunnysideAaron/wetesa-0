package server

import (
	"api/internal/database"
	"log"
	"net/http"
)

// addRoutes maps all the API routes
func addRoutes(mux *http.ServeMux, logger *log.Logger, db *database.Postgres) {
	mux.Handle("GET /api/clients", handleListClients(logger, db))
	mux.Handle("GET /api/clients/{id}", handleGetClient(logger, db))
	mux.Handle("POST /api/clients", handleCreateClient(logger, db))
	mux.Handle("PUT /api/clients/{id}", handleUpdateClient(logger, db))
	mux.Handle("DELETE /api/clients/{id}", handleDeleteClient(logger, db))
	mux.Handle("GET /healthz", handleHealthz())
	mux.Handle("GET /healthdbz", handleHealthDBz(logger, db))
	mux.Handle("/", http.NotFoundHandler())
}
