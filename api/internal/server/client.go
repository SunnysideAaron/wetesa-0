package server

import (
	"api/internal/database"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// handleListClients handles requests to list all clients
func handleListClients(logger *log.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clients, err := db.GetClients(r.Context())
		if err != nil {
			logger.Printf("error getting clients: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if err := encode(w, r, http.StatusOK, clients); err != nil {
			logger.Printf("error encoding response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}

// handleGetUser handles requests to get a specific user
func handleGetClient(logger *log.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Missing ID", http.StatusBadRequest)
			return
		}

		client, err := db.GetClient(r.Context(), id)
		if err != nil {
			logger.Printf("error getting client: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if err := encode(w, r, http.StatusOK, client); err != nil {
			logger.Printf("error encoding response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}

// handleCreateClient handles requests to create a new client
func handleCreateClient(logger *log.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := database.Client{}

		if err := decode(r, &client); err != nil {
			logger.Printf("error decoding request: %v", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			encode(w, r, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if problems := client.Valid(r.Context()); len(problems) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			encode(w, r, http.StatusBadRequest, problems)
			return
		}

		err := db.InsertClient(r.Context(), client)
		if err != nil {
			logger.Printf("error creating client: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Return the client data
		response := map[string]string{
			"name":    client.Name,
			"address": client.Address.String,
		}

		if err := encode(w, r, http.StatusCreated, response); err != nil {
			logger.Printf("error encoding response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}

// handleUpdateClient handles requests to update an existing client
func handleUpdateClient(logger *log.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Missing ID", http.StatusBadRequest)
			return
		}

		// First get the existing client
		_, err := db.GetClient(r.Context(), id)
		if err != nil {
			logger.Printf("error getting client: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Decode the update request
		var updateClient database.Client
		if err := decode(r, &updateClient); err != nil {
			logger.Printf("error decoding request: %v", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			encode(w, r, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		clientID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		updateClient.Client_id = clientID

		// Validate the update request
		if problems := updateClient.Valid(r.Context()); len(problems) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			encode(w, r, http.StatusBadRequest, problems)
			return
		}

		// Perform the update
		err = db.UpdateClient(r.Context(), updateClient)
		if err != nil {
			logger.Printf("error updating client: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Return the updated client
		response := map[string]string{
			"name":    updateClient.Name,
			"address": updateClient.Address.String,
		}

		if err := encode(w, r, http.StatusOK, response); err != nil {
			logger.Printf("error encoding response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}

// handleDeleteClient handles requests to delete a client
func handleDeleteClient(logger *log.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			logger.Printf("error deleting client: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Return 204 No Content for successful deletion
		w.WriteHeader(http.StatusNoContent)
	})
}
