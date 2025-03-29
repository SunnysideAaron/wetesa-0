package server

import (
	"api/internal/database"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

// handleListClients handles requests to list all clients
func handleListClients(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			clients, err := db.GetClients(r.Context())
			if err != nil {
				logger.Error("error getting clients", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if err := encode(w, r, http.StatusOK, clients); err != nil {
				logger.Error("error encoding response", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		},
	)
}

// handleGetClient handles requests to get a specific client
func handleGetClient(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")
			if id == "" {
				http.Error(w, "Missing ID", http.StatusBadRequest)
				return
			}

			client, err := db.GetClient(r.Context(), id)
			if err != nil {
				logger.Error("error getting client", "error", err, "id", id)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if err := encode(w, r, http.StatusOK, client); err != nil {
				logger.Error("error encoding response", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// example of logging while hiding sensitive information
			logger.LogAttrs(
				r.Context(),
				slog.LevelDebug,
				"client retrieved",
				slog.Any("client", client.LogValue()),
			)

		},
	)
}

// handleCreateClient handles requests to create a new client
func handleCreateClient(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			client, problems, err := decode[database.Client](r)
			if err != nil {
				logger.Error("error decoding request", "error", err, "problems", problems)
				if err := encode(w, r, http.StatusBadRequest, map[string]interface{}{
					"error":    err.Error(),
					"problems": problems,
				}); err != nil {
					logger.Error("error encoding response", "error", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
				return
			}

			err = db.InsertClient(r.Context(), client)
			if err != nil {
				logger.Error("error creating client", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Return the client data
			response := map[string]string{
				"name":    client.Name,
				"address": client.Address.String,
			}

			if err := encode(w, r, http.StatusCreated, response); err != nil {
				logger.Error("error encoding response", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		},
	)
}

// handleUpdateClient handles requests to update an existing client
// TODO this is for a PUT request. Which is OK but we might want to use PATCH instead.
func handleUpdateClient(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")
			if id == "" {
				http.Error(w, "Missing ID", http.StatusBadRequest)
				return
			}

			// First get the existing client
			_, err := db.GetClient(r.Context(), id)
			if err != nil {
				logger.Error("error getting client", "error", err, "id", id)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Decode the update request
			updateClient, problems, err := decode[database.Client](r)
			if err != nil {
				logger.Error("error decoding request", "error", err, "problems", problems)
				if err := encode(w, r, http.StatusBadRequest, map[string]interface{}{
					"error":    err.Error(),
					"problems": problems,
				}); err != nil {
					logger.Error("error encoding response", "error", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
				return
			}

			clientID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "Invalid ID format", http.StatusBadRequest)
				return
			}

			updateClient.Client_id = clientID

			// Perform the update
			err = db.UpdateClient(r.Context(), updateClient)
			if err != nil {
				logger.Error("error updating client", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Return the updated client
			response := map[string]string{
				"name":    updateClient.Name,
				"address": updateClient.Address.String,
			}

			if err := encode(w, r, http.StatusOK, response); err != nil {
				logger.Error("error encoding response", "error", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		},
	)
}

// handleDeleteClient handles requests to delete a client
func handleDeleteClient(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")

			// TODO research i don't think it's possible to get to this method without an id
			if id == "" {
				http.Error(w, "Missing ID", http.StatusBadRequest)
				return
			}

			err := db.DeleteClient(r.Context(), id)
			if err != nil {
				if strings.Contains(err.Error(), "not found") {
					http.Error(w, "Client not found", http.StatusNotFound)
					return
				}
				logger.Error("error deleting client", "error", err, "id", id)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Return 204 No Content for successful deletion
			w.WriteHeader(http.StatusNoContent)
		},
	)
}
