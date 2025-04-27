package database

import (
	"context"
	"log/slog"
	"strconv"
)

// Health performs a database health check and returns status information.
// It returns a map containing connection pool statistics and status indicators.
func (pg *Postgres) Health(ctx context.Context, logger *slog.Logger) map[string]string {
	stats := make(map[string]string)

	// Ping the database
	err := pg.pool.Ping(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = "db down"

		// err contains some sensitive information. don't show users.
		// TODO do we even want to log these values? or is that a security hole?
		logger.LogAttrs(
			ctx,
			slog.LevelWarn, // Perhaps db will come back up. Warning for now. If stays down that is an error.
			"db down",
			slog.String("error", err.Error()),
			// slog.Any("error", err), // TODO is slog.Any properly handled in PrettyHandler? For some reason log doesn't spit out.
		)

		return stats
	}

	stats["status"] = "up"
	stats["message"] = "It's healthy"

	dbStats := pg.pool.Stat()

	stats["AcquireCount"] = strconv.FormatInt(dbStats.AcquireCount(), 10)
	stats["AcquireDuration"] = strconv.FormatInt(int64(dbStats.AcquireDuration()), 10)
	stats["AcquiredConns"] = strconv.Itoa(int(dbStats.AcquiredConns()))
	stats["CanceledAcquireCount"] = strconv.FormatInt(dbStats.CanceledAcquireCount(), 10)
	stats["ConstructingConns"] = strconv.Itoa(int(dbStats.ConstructingConns()))
	stats["EmptyAcquireCount"] = strconv.FormatInt(dbStats.EmptyAcquireCount(), 10)
	stats["IdleConns"] = strconv.Itoa(int(dbStats.IdleConns()))
	stats["MaxConns"] = strconv.Itoa(int(dbStats.MaxConns()))
	stats["MaxIdleDestroyCount"] = strconv.FormatInt(dbStats.MaxIdleDestroyCount(), 10)
	stats["MaxLifetimeDestroyCount"] = strconv.FormatInt(dbStats.MaxLifetimeDestroyCount(), 10)
	stats["NewConnsCount"] = strconv.FormatInt(dbStats.NewConnsCount(), 10)
	stats["TotalConns"] = strconv.Itoa(int(dbStats.TotalConns()))

	return stats
}
