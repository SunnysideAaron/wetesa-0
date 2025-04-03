package database

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

// TODO convert to error constants and types. rest of code as well.

// Client represents a client record in the database.
// It contains basic information about a client including their unique identifier,
// name, and optional address.
type Client struct {
	ClientID int         `json:"client_id"`
	Name     string      `json:"name"`
	Address  pgtype.Text `json:"address"`
}

// Valid implements the Validator interface for Client.
// It checks if the required fields are properly set and returns a map of validation errors.
// An empty map indicates the Client is valid.
func (c Client) Valid(_ context.Context) map[string]string {
	problems := make(map[string]string)

	if c.Name == "" {
		problems["name"] = "name is required"
	}

	// Address is optional, so no validation needed

	return problems
}

// LogValue implements slog.LogValuer to provide structured logging support.
// It returns the client's ID as the log value for concise logging.
func (c Client) LogValue() slog.Value {
	return slog.IntValue(c.ClientID)
}

// InsertClient adds a new client record to the database.
// It returns an error if the insert operation fails.
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

// BulkInsertClients is slower than CopyInserts. Use bulk inserts if you need to know a particular insert failed.
func (pg *Postgres) BulkInsertClients(ctx context.Context, clients []Client) error {
	query := `INSERT INTO client (name, address) VALUES (@name, @address)`

	batch := &pgx.Batch{} //nolint:exhaustruct // works fine. we don't need to initialize. already handled.

	for _, client := range clients {
		args := pgx.NamedArgs{
			"name":    client.Name,
			"address": client.Address,
		}
		batch.Queue(query, args)
	}

	results := pg.pool.SendBatch(ctx, batch)
	// defer results.Close()
	defer func() {
		err := results.Close()
		if err != nil {
			slog.LogAttrs(
				ctx,
				slog.LevelError,
				"could not close results",
				slog.String("error", err.Error()),
			)
		}
	}()

	for _, client := range clients {
		_, err := results.Exec()
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
				slog.LogAttrs(
					ctx,
					slog.LevelInfo,
					"user already exists",
					slog.String("name", client.Name),
				)
				continue
			}

			fmt.Println("unable to insert row: %w", err)
		}
	}

	return results.Close()
}

// CopyInsertClients if faster than BulkInsertClients
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
		return fmt.Errorf("error copying into %s table: %w", tableName, err)
	}

	return nil
}

// GetClients retrieves a list of clients from the database.
// TODO pagination.
func (pg *Postgres) GetClients(ctx context.Context) ([]Client, error) {
	query := `SELECT client_id, name, address FROM client order by client_id desc LIMIT 10`

	rows, err := pg.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query clients: %w", err)
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Client])
}

// GetClient retrieves a single client by their ID.
// Returns the client data if found, or an error if the client doesn't exist
// or if the query fails.
func (pg *Postgres) GetClient(ctx context.Context, id string) (Client, error) {
	var client Client

	query := `SELECT client_id, name, address FROM client WHERE client_id = $1`

	row := pg.pool.QueryRow(ctx, query, id)

	err := row.Scan(&client.ClientID, &client.Name, &client.Address)
	if err != nil {
		return client, err
	}
	return client, nil
}

// UpdateClient updates an existing client's information.
// Returns an error if the client doesn't exist or if the update operation fails.
// TODO this is for a PUT request. Which is OK but we might want to use PATCH instead.
func (pg *Postgres) UpdateClient(ctx context.Context, c Client) error {
	query := `UPDATE client 
			  SET name = @name, address = @address 
			  WHERE client_id = @client_id`

	args := pgx.NamedArgs{
		"client_id": c.ClientID,
		"name":      c.Name,
		"address":   c.Address,
	}

	result, err := pg.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update client: %w", err)
	}

	// Check if any row was actually updated
	if result.RowsAffected() == 0 {
		return fmt.Errorf("client with id %d not found", c.ClientID)
	}

	return nil
}

// DeleteClient removes a client from the database by their ID.
// Returns an error if the client doesn't exist or if the deletion fails.
func (pg *Postgres) DeleteClient(ctx context.Context, id string) error {
	query := `DELETE FROM client WHERE client_id = $1`

	result, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("unable to delete client: %w", err)
	}

	// Check if any row was actually deleted
	if result.RowsAffected() == 0 {
		return fmt.Errorf("client with id %s not found", id)
	}

	return nil
}
