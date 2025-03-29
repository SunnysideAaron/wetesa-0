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
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context, pCfg *pgxpool.Config, logger *slog.Logger) *Postgres {
	pgOnce.Do(func() {
		pg, err := pgxpool.NewWithConfig(ctx, pCfg)
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

func (pg *Postgres) Close() {
	pg.pool.Close()
}
