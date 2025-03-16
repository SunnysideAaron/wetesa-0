// Generated from poe using cluade

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	//"github.com/jackc/pgx/v5/pgtype"

	"api/internal/database"
)

// NewServer creates a new HTTP server
func NewServer(logger *log.Logger, db *database.Postgres) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger, db)

	var handler http.Handler = mux
	handler = loggingMiddleware(logger, handler)

	return handler
}

// addRoutes maps all the API routes
func addRoutes(mux *http.ServeMux, logger *log.Logger, db *database.Postgres) {
	mux.Handle("GET /api/clients", handleListClients(logger, db))
	mux.Handle("GET /api/clients/{id}", handleGetClient(logger, db))
	mux.Handle("POST /api/clients", handleCreateClient(logger, db))
	mux.Handle("PUT /api/clients/{id}", handleUpdateClient(logger, db))
	mux.Handle("DELETE /api/clients/{id}", handleDeleteClient(logger, db))
	mux.Handle("GET /healthz", handleHealthz())
	mux.Handle("/", http.NotFoundHandler())
}

//TODO middleware and middleware grouping

// loggingMiddleware is middleware that logs requests
func loggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// handleHealthz handles health check requests
func handleHealthz() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}

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

// TODO why is encode() and decode() different than article as written
// encode encodes the response as JSON
func encode(w http.ResponseWriter, r *http.Request, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// decode decodes the request body
func decode(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("reading body: %w", err)
	}

	// Log the received body
	log.Printf("Received body: %s", string(body))

	// Check for common JSON formatting issues
	if len(body) == 0 {
		return fmt.Errorf("empty request body")
	}

	if body[0] == '\'' {
		return fmt.Errorf("invalid JSON: use double quotes (\") instead of single quotes (')")
	}

	if !strings.Contains(string(body), "\"") {
		return fmt.Errorf("invalid JSON: property names and string values must be enclosed in double quotes")
	}

	// Try to decode
	err = json.NewDecoder(bytes.NewReader(body)).Decode(v)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "looking for beginning of value"):
			return fmt.Errorf("invalid JSON format: please check your JSON syntax")
		case strings.Contains(err.Error(), "cannot unmarshal"):
			return fmt.Errorf("invalid data type in JSON: %v", err)
		default:
			return fmt.Errorf("JSON decode error: %v", err)
		}
	}

	return nil
}

// run is the actual main function
func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := log.New(w, "", log.LstdFlags)

	// Create database connection
	//db, err := database.NewPG(ctx)

	db := database.NewPG(ctx)

	// if err != nil {
	// 	return fmt.Errorf("failed to connect to database: %w", err)
	// }
	defer db.Close()

	// Create a new server
	srv := NewServer(logger, db)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	// Configure the HTTP server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: srv,
	}

	// Start the server in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		logger.Printf("server listening on %s", httpServer.Addr)
		serverErrors <- httpServer.ListenAndServe()
	}()

	// Wait for interrupt or error
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		logger.Println("shutting down server...")

		// Create shutdown context with timeout
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		// Attempt graceful shutdown
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server shutdown error: %w", err)
		}
	}

	return nil
}

func main() {
	if err := run(context.Background(), os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
