package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"api/internal/database"
)

type Server struct {
	port int
	pg   *database.Postgres
}

func NewServer(ctx context.Context) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("API_PORT"))
	NewServer := &Server{
		port: port,
		pg:   database.NewPG(ctx),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
