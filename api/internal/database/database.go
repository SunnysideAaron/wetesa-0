// Informative guides.
// https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
// https://hexacluster.ai/postgresql/connecting-to-postgresql-with-go-using-pgx/
package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func readConfig() (*pgxpool.Config, error) {
	// This example worked to set config values.
	// https://github.com/jackc/pgx/issues/588
	config, _ := pgxpool.ParseConfig("")
	config.ConnConfig.Host = os.Getenv("DATASTORE_HOST")
	port, err := strconv.Atoi(os.Getenv("DATASTORE_PORT"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid port number:", err)
		os.Exit(1)
	}
	config.ConnConfig.Port = uint16(port)
	config.ConnConfig.Database = os.Getenv("POSTGRESQL_DATABASE")
	config.ConnConfig.User = os.Getenv("POSTGRESQL_USERNAME")
	config.ConnConfig.Password = os.Getenv("POSTGRESQL_PASSWORD")
	// TODO pool settings should at least come from env vars.
	// can they be changed while the api is running? If so then it
	//  should be possible to do so for troubleshooting. purposes
	// is that true of the db settings as well?
	// example pool config TODO figure best values for this application.
	config.MaxConns = 5

	return config, nil
}

func NewPG(ctx context.Context) *Postgres {

	config, err := readConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read config: ", err)
		os.Exit(1)
	}

	pgOnce.Do(func() {
		pg, pgErr := pgxpool.NewWithConfig(ctx, config)
		if pgErr != nil {
			fmt.Fprintln(os.Stderr, "unable to create connection pool: %w", pgErr)
			return
		}

		pgInstance = &Postgres{pg}
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to create connection pool2: %w", err)
		return nil
	}

	return pgInstance
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.pool.Ping(ctx)
}

// https://pkg.go.dev/github.com/jackc/pgx/v5@v5.7.2/pgxpool#Stat
// TODO health check

func (pg *Postgres) Close() {
	pg.pool.Close()
}
