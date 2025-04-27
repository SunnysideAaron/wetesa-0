package database

import (
	"context"
	"log/slog"
)

// HealthStatus represents the health check response data
type HealthStatus struct {
	Status                  string `json:"status"`
	Message                 string `json:"message,omitempty"`
	Error                   string `json:"error,omitempty"`
	AcquireCount            int64  `json:"acquire_count"`
	AcquireDuration         int64  `json:"acquire_duration"`
	AcquiredConns           int    `json:"acquired_conns"`
	CanceledAcquireCount    int64  `json:"canceled_acquire_count"`
	ConstructingConns       int    `json:"constructing_conns"`
	EmptyAcquireCount       int64  `json:"empty_acquire_count"`
	IdleConns               int    `json:"idle_conns"`
	MaxConns                int    `json:"max_conns"`
	MaxIdleDestroyCount     int64  `json:"max_idle_destroy_count"`
	MaxLifetimeDestroyCount int64  `json:"max_lifetime_destroy_count"`
	NewConnsCount           int64  `json:"new_conns_count"`
	TotalConns              int    `json:"total_conns"`
}

// Health performs a database health check and returns status information.
// It returns a HealthStatus containing connection pool statistics and status indicators.
func (pg *Postgres) Health(ctx context.Context, logger *slog.Logger) HealthStatus {
	var status HealthStatus

	// Ping the database
	err := pg.pool.Ping(ctx)
	if err != nil {
		status.Status = "down"
		status.Error = "db down"

		// err contains some sensitive information. don't show users.
		// TODO do we even want to log these values? or is that a security hole?
		logger.LogAttrs(
			ctx,
			slog.LevelWarn, // Perhaps db will come back up. Warning for now. If stays down that is an error.
			"db down",
			slog.String("error", err.Error()),
			// slog.Any("error", err), // TODO is slog.Any properly handled in PrettyHandler? For some reason log doesn't spit out.
		)

		return status
	}

	status.Status = "up"
	status.Message = "It's healthy"

	dbStats := pg.pool.Stat()

	status.AcquireCount = dbStats.AcquireCount()
	status.AcquireDuration = int64(dbStats.AcquireDuration())
	status.AcquiredConns = int(dbStats.AcquiredConns())
	status.CanceledAcquireCount = dbStats.CanceledAcquireCount()
	status.ConstructingConns = int(dbStats.ConstructingConns())
	status.EmptyAcquireCount = dbStats.EmptyAcquireCount()
	status.IdleConns = int(dbStats.IdleConns())
	status.MaxConns = int(dbStats.MaxConns())
	status.MaxIdleDestroyCount = dbStats.MaxIdleDestroyCount()
	status.MaxLifetimeDestroyCount = dbStats.MaxLifetimeDestroyCount()
	status.NewConnsCount = dbStats.NewConnsCount()
	status.TotalConns = int(dbStats.TotalConns())

	return status
}
