package database

import (
	"context"
	"log/slog"
	"strconv"
)

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (pg *Postgres) Health(ctx context.Context, logger *slog.Logger) map[string]string {
	stats := make(map[string]string)

	// Ping the database
	err := pg.pool.Ping(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = "db down"

		// err contains some sensitive information. don't show users.
		// TOOD do we even want to log these values? or is that a security hole?
		logger.LogAttrs(
			ctx,
			slog.LevelWarn, //Perhaps db will come back up. Warning for now. If stays down that is an error.
			"db down",
			slog.String("error", err.Error()),
			//slog.Any("error", err), // TODO is slog.Any properly handled in PrettyHandler? For some reason log doesn't spit out.
		)

		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
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

	// stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	// stats["in_use"] = strconv.Itoa(dbStats.InUse)
	// stats["idle"] = strconv.Itoa(dbStats.Idle)
	// stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	// stats["wait_duration"] = dbStats.WaitDuration.String()
	// stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	// stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	//TODO move the limits to config.
	//TODO don't overwrite message. append to message
	// Evaluate stats to provide a health message
	// if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
	// 	stats["message"] = "The database is experiencing heavy load."
	// }

	// if dbStats.WaitCount > 1000 {
	// 	stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	// }

	// if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
	// 	stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	// }

	// if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
	// 	stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	// }

	return stats
}
