package server

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"api/internal/database"
)

// handleListClients handles requests to list all clients
func handleListClients(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Default values for pagination and sorting
			page := 0
			size := 10
			sort := "asc"

			// Parse page from query parameter
			if pageStr := r.URL.Query().Get("page"); pageStr != "" {
				parsedPage, err := strconv.Atoi(pageStr)
				if err != nil {
					http.Error(w, "Invalid page parameter", http.StatusBadRequest)
					return
				}
				if parsedPage < 0 {
					parsedPage = 0
				}
				page = parsedPage
			}

			// Parse size from query parameter
			if sizeStr := r.URL.Query().Get("size"); sizeStr != "" {
				parsedSize, err := strconv.Atoi(sizeStr)
				if err != nil {
					http.Error(w, "Invalid size parameter", http.StatusBadRequest)
					return
				}
				// Enforce reasonable size limits
				if parsedSize > 100 {
					parsedSize = 100
				} else if parsedSize < 1 {
					parsedSize = 1
				}
				size = parsedSize
			}

			// Parse sort from query parameter
			if sortStr := r.URL.Query().Get("sort"); sortStr != "" {
				sortLower := strings.ToLower(sortStr)
				if sortLower != "asc" && sortLower != "desc" {
					http.Error(w, "Invalid sort parameter. Must be 'asc' or 'desc'", http.StatusBadRequest)
					return
				}

				sort = sortLower
			}

			// Get filter parameters
			filters := database.ClientFilters{
				Name:    r.URL.Query().Get("name"),
				Address: r.URL.Query().Get("address"),
			}

			// Calculate offset from page and size
			offset := page * size

			clients, err := db.GetClients(r.Context(), size, offset, sort, filters)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error getting clients",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response := map[string]any{
				"clients":  clients,
				"page":     page,
				"size":     size,
				"sort":     sort,
				"filters":  filters,
				"returned": len(clients),
			}

			err = encode(w, r, http.StatusOK, response)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error encoding response",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
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
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error getting client",
					slog.String("error", err.Error()),
					slog.String("id", id),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = encode(w, r, http.StatusOK, client)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error encoding response",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error decoding request",
					slog.String("error", err.Error()),
					slog.Any("problems", problems),
				)
				if err = encode(w, r, http.StatusBadRequest, map[string]any{
					"error":    err.Error(),
					"problems": problems,
				}); err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelInfo,
						"error encoding response",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			err = db.InsertClient(r.Context(), client)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error creating client",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Return the client data
			response := map[string]string{
				"name":    client.Name,
				"address": client.Address.String,
			}

			err = encode(w, r, http.StatusCreated, response)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error encoding response",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error getting client",
					slog.String("error", err.Error()),
					slog.String("id", id),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Decode the update request
			updateClient, problems, err := decode[database.Client](r)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error decoding request",
					slog.String("error", err.Error()),
					slog.Any("problems", problems),
				)
				if err = encode(w, r, http.StatusBadRequest, map[string]any{
					"error":    err.Error(),
					"problems": problems,
				}); err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelInfo,
						"error encoding response",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			updateClient.ClientID = id

			// Perform the update
			err = db.UpdateClient(r.Context(), updateClient)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error updating client",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Return the updated client
			response := map[string]string{
				"name":    updateClient.Name,
				"address": updateClient.Address.String,
			}

			err = encode(w, r, http.StatusOK, response)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error encoding response",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error deleting client",
					slog.String("error", err.Error()),
					slog.String("id", id),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Return 204 No Content for successful deletion
			w.WriteHeader(http.StatusNoContent)
		},
	)
}
