// Following guidelines from:
// [How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"api/internal/config"
	"api/internal/database"
	"api/internal/server"
)

// run is the actual main function
// [func main() only calls run()](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#func-main-only-calls-run)
func run(ctx context.Context, w io.Writer, cfg *config.APIConfig) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := log.New(w, "", log.LstdFlags)

	// Create database connection
	db := database.NewPG(ctx)
	defer db.Close()

	// Create a new server
	srv := server.NewServer(logger, db)

	// Configure the HTTP server
	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort),
		Handler:      srv,
		ReadTimeout:  cfg.APIReadTimeout * time.Second,
		WriteTimeout: cfg.APIWriteTimeout * time.Second,
		IdleTimeout:  cfg.APIIdleTimeout * time.Second,
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
		// [Gracefully shutting down](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#gracefully-shutting-down)

		logger.Println("shutting down server...")

		// Create shutdown context with timeout
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		// Attempt graceful shutdown
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server shutdown error: %w", err)
		}
	}

	return nil
}

func main() {
	ctx := context.Background()
	cfg := config.LoadAPIConfig()

	// TODO break all these single line error check statements into multiple lines.
	if err := run(ctx, os.Stdout, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
