// Following guidelines from:
// [How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"api/internal/config"
	"api/internal/database"
	"api/internal/logging"
	"api/internal/server"

	"github.com/jackc/pgx/v5/pgxpool"
)

// run is the actual main function
// [func main() only calls run()](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#func-main-only-calls-run)
func run(
	ctx context.Context,
	cfg *config.APIConfig,
	pCfg *pgxpool.Config,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := logging.NewLogger(cfg)
	slog.SetDefault(logger)

	// Create database connection
	db := database.NewPG(ctx, pCfg)
	defer db.Close()

	handle := server.AddRoutes(logger, cfg, db)

	// Configure the HTTP server
	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort),
		Handler:      handle,
		ReadTimeout:  cfg.APIReadTimeout * time.Second,
		WriteTimeout: cfg.APIWriteTimeout * time.Second,
		IdleTimeout:  cfg.APIIdleTimeout * time.Second,
	}

	// Start the server in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("server starting", "addr", httpServer.Addr)
		serverErrors <- httpServer.ListenAndServe()
	}()

	// Wait for interrupt or error
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		// [Gracefully shutting down](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#gracefully-shutting-down)

		logger.Info("shutting down server...")

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
	pCfg := config.LoadDBConfig()

	// TODO break all these single line error check statements into multiple lines.
	// finder linter and formatter
	if err := run(ctx, cfg, pCfg); err != nil {
		slog.Error("application error", "error", err)
		os.Exit(1)
	}
}
