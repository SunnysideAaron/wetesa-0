// Generated from poe using cluade

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

// User represents a user in our system
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// UserStore is a simple in-memory store for users
type UserStore struct {
	sync.RWMutex
	users map[string]User
}

// NewUserStore creates a new user store
func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]User),
	}
}

// Get returns a user by ID
func (s *UserStore) Get(id string) (User, bool) {
	s.RLock()
	defer s.RUnlock()
	user, ok := s.users[id]
	return user, ok
}

// List returns all users
func (s *UserStore) List() []User {
	s.RLock()
	defer s.RUnlock()
	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// Create adds a new user
func (s *UserStore) Create(user User) {
	s.Lock()
	defer s.Unlock()
	s.users[user.ID] = user
}

// Update modifies an existing user
func (s *UserStore) Update(user User) bool {
	s.Lock()
	defer s.Unlock()
	_, exists := s.users[user.ID]
	if !exists {
		return false
	}
	s.users[user.ID] = user
	return true
}

// Delete removes a user
func (s *UserStore) Delete(id string) bool {
	s.Lock()
	defer s.Unlock()
	_, exists := s.users[id]
	if !exists {
		return false
	}
	delete(s.users, id)
	return true
}

// CreateUserRequest is the request body for creating a user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Valid validates the create user request
func (r CreateUserRequest) Valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	if r.Name == "" {
		problems["name"] = "name is required"
	}

	if r.Email == "" {
		problems["email"] = "email is required"
	} else if !strings.Contains(r.Email, "@") {
		problems["email"] = "email must be valid"
	}

	return problems
}

// UpdateUserRequest is the request body for updating a user
type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Valid validates the update user request
func (r UpdateUserRequest) Valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	if r.Email != "" && !strings.Contains(r.Email, "@") {
		problems["email"] = "email must be valid"
	}

	return problems
}

// Validator is an object that can be validated
type Validator interface {
	Valid(ctx context.Context) map[string]string
}

// NewServer creates a new HTTP server
func NewServer(logger *log.Logger, userStore *UserStore) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger, userStore)

	var handler http.Handler = mux
	handler = loggingMiddleware(logger, handler)

	return handler
}

// addRoutes maps all the API routes
func addRoutes(mux *http.ServeMux, logger *log.Logger, userStore *UserStore) {
	mux.Handle("GET /api/users", handleListUsers(logger, userStore))
	mux.Handle("GET /api/users/{id}", handleGetUser(logger, userStore))
	mux.Handle("POST /api/users", handleCreateUser(logger, userStore))
	mux.Handle("PUT /api/users/{id}", handleUpdateUser(logger, userStore))
	mux.Handle("DELETE /api/users/{id}", handleDeleteUser(logger, userStore))
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

// handleListUsers handles requests to list all users
func handleListUsers(logger *log.Logger, store *UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users := store.List()
		if err := encode(w, r, http.StatusOK, users); err != nil {
			logger.Printf("error encoding response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}

// handleGetUser handles requests to get a specific user
func handleGetUser(logger *log.Logger, store *UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Missing ID", http.StatusBadRequest)
			return
		}

		user, found := store.Get(id)
		if !found {
			http.NotFound(w, r)
			return
		}

		if err := encode(w, r, http.StatusOK, user); err != nil {
			logger.Printf("error encoding response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}

// handleCreateUser handles requests to create a new user
func handleCreateUser(logger *log.Logger, store *UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req CreateUserRequest
		if err := decode(r, &req); err != nil {
			logger.Printf("error decoding request: %v", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			encode(w, r, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if problems := req.Valid(r.Context()); len(problems) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			encode(w, r, http.StatusBadRequest, problems)
			return
		}

		user := User{
			ID:        fmt.Sprintf("usr_%d", time.Now().UnixNano()),
			Name:      req.Name,
			Email:     req.Email,
			CreatedAt: time.Now(),
		}

		store.Create(user)

		if err := encode(w, r, http.StatusCreated, user); err != nil {
			logger.Printf("error encoding response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}

// handleUpdateUser handles requests to update an existing user
func handleUpdateUser(logger *log.Logger, store *UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Missing ID", http.StatusBadRequest)
			return
		}

		user, found := store.Get(id)
		if !found {
			http.NotFound(w, r)
			return
		}

		var req UpdateUserRequest
		if err := decode(r, &req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if problems := req.Valid(r.Context()); len(problems) > 0 {
			encode(w, r, http.StatusBadRequest, problems)
			return
		}

		// Update fields if provided
		if req.Name != "" {
			user.Name = req.Name
		}
		if req.Email != "" {
			user.Email = req.Email
		}

		if success := store.Update(user); !success {
			http.NotFound(w, r)
			return
		}

		if err := encode(w, r, http.StatusOK, user); err != nil {
			logger.Printf("error encoding response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}

// handleDeleteUser handles requests to delete a user
func handleDeleteUser(logger *log.Logger, store *UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Missing ID", http.StatusBadRequest)
			return
		}

		if success := store.Delete(id); !success {
			http.NotFound(w, r)
			return
		}

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

	// Create a new user store
	userStore := NewUserStore()

	// Create a new server
	srv := NewServer(logger, userStore)

	// Configure the HTTP server
	httpServer := &http.Server{
		Addr: net.JoinHostPort("0.0.0.0", "8080"),
		//Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: srv,
	}

	// Start the server in a goroutine
	// TODO should this have context?
	// how do handlers know context?
	go func() {
		logger.Printf("server listening on %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	// Wait for interrupt signal
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		logger.Println("shutting down server...")

		// Create a deadline to wait for
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// This implementation follows Mat Ryer's guidelines:

//     NewServer constructor pattern: The server setup is handled in a constructor function that takes dependencies and returns an http.Handler.

//     Routes mapping in one place: All API routes are defined in the addRoutes function for clarity.

//     main() only calls run(): The main function just calls the run function, which does all the work and returns an error.

//     Maker funcs return the handler: All handler functions like handleListUsers return http.Handler types.

//     Validation interface: Uses a Validator interface with a Valid method that returns a map of problems.

//     Middleware pattern: Implements middleware as functions that take an http.Handler and return a new one.

//     Graceful shutdown: Properly handles shutdown signals and gives running requests time to complete.

//     Encode/decode helpers: Uses helper functions to handle JSON encoding/decoding in one place.

// The example provides a complete CRUD API for managing users with:

//     List all users (GET /api/users)
//     Get a single user (GET /api/users/{id})
//     Create a user (POST /api/users)
//     Update a user (PUT /api/users/{id})
//     Delete a user (DELETE /api/users/{id})

// Would you like me to explain any specific part of this implementation in more detail?
