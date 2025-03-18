// Informative guides.
// https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
// https://hexacluster.ai/postgresql/connecting-to-postgresql-with-go-using-pgx/
package database

import (
	"context"
	"fmt"
	"os"
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

func NewPG(ctx context.Context, pCfg *pgxpool.Config) *Postgres {
	pgOnce.Do(func() {
		pg, err := pgxpool.NewWithConfig(ctx, pCfg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to create connection pool: %w", err)
			return
		}

		pgInstance = &Postgres{pg}
	})

	return pgInstance
}

func (pg *Postgres) Close() {
	pg.pool.Close()
}
