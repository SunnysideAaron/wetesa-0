// Package database provides database connectivity and operations for the API service.
//
// Informative guides.
// https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
// https://hexacluster.ai/postgresql/connecting-to-postgresql-with-go-using-pgx/
package database

import (
	"context"
	"log/slog"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres represents a PostgreSQL database connection pool and related configuration.
type Postgres struct {
	pool *pgxpool.Pool
}

// Validator is an object that can be validated.
type Validator interface {
	// Valid checks the object and returns any
	// problems. If len(problems) == 0 then
	// the object is valid.
	Valid(ctx context.Context) (problems map[string]string)
}

var (
	pgInstance *Postgres //nolint:gochecknoglobals // singleton pattern for database connection
	pgOnce     sync.Once //nolint:gochecknoglobals // singleton pattern for database connection
)

// NewPG creates and initializes a new PostgreSQL database connection pool.
// It verifies the connection and returns an error if the connection fails.
func NewPG(ctx context.Context, cfg *pgxpool.Config, logger *slog.Logger) *Postgres {
	pgOnce.Do(func() {
		pg, err := pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			logger.LogAttrs(
				ctx,
				slog.LevelError,
				"unable to create connection pool",
				slog.String("error", err.Error()),
			)
			return
		}

		pgInstance = &Postgres{pg}
	})

	return pgInstance
}

// Close releases all connections in the pool and shuts down the connection pool.
// This method should be called when the application is shutting down to ensure
// proper cleanup of database resources.
func (pg *Postgres) Close() {
	pg.pool.Close()
}
