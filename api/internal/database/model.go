package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Client struct {
	Client_id int         `json:"client_id"`
	Name      string      `json:"name"`
	Address   pgtype.Text `json:"address"`
}

func (pg *Postgres) InsertClient(ctx context.Context, c Client) error {
	query := `INSERT INTO client (name, address) VALUES (@name, @address)`
	args := pgx.NamedArgs{
		"name":    c.Name,
		"address": c.Address,
	}

	_, err := pg.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

// CopyInserts is faster. Use bulk inserts if you need to know a particular insert failed.
func (pg *Postgres) BulkInsertClients(ctx context.Context, clients []Client) error {
	query := `INSERT INTO client (name, address) VALUES (@name, @address)`

	batch := &pgx.Batch{}
	for _, client := range clients {
		args := pgx.NamedArgs{
			"name":    client.Name,
			"address": client.Address,
		}
		batch.Queue(query, args)
	}

	results := pg.pool.SendBatch(ctx, batch)
	defer results.Close()

	for _, client := range clients {
		_, err := results.Exec()
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
				log.Printf("user %s already exists", client.Name)
				continue
			}

			fmt.Println("unable to insert row: %w", err)
		}
	}

	return results.Close()
}

// See note on BulkInsertClients
func (pg *Postgres) CopyInsertClients(ctx context.Context, clients []Client) error {
	entries := [][]any{}
	columns := []string{"name", "address"}
	tableName := "client"

	for _, client := range clients {
		entries = append(entries, []any{client.Name, client.Address})
	}

	_, err := pg.pool.CopyFrom(
		ctx,
		pgx.Identifier{tableName},
		columns,
		pgx.CopyFromRows(entries),
	)

	if err != nil {
		return fmt.Errorf("error copying into %s table: %v", tableName, err)
	}

	return nil
}

// TODO pagination.
func (pg *Postgres) GetClients(ctx context.Context) ([]Client, error) {
	query := `SELECT client_id, name, address FROM client order by client_id desc LIMIT 10`

	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Client])
}
