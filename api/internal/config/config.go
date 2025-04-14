// Package config provides configuration management for the API service.
// [The Twelve-Factor App III. Config](https://12factor.net/config)
package config

import (
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	// EnvironmentDev represents the development environment configuration.
	EnvironmentDev string = "dev"
	// EnvironmentProd represents the production environment configuration.
	EnvironmentProd string = "prod"
)

// APIConfig stores the API service configuration parameters.
type APIConfig struct {
	// BaseURL is the base URL for the API. Includes API version. Unlike the
	// other settings do not load BaseURL from an environment variable. Including
	// here so there is only one source of truth for api version.
	BaseURL string
	// Environment specifies the running environment (dev/prod)
	Environment string
	// APIHost is the host address to bind the server to
	APIHost string
	// APIPort is the port number to listen on
	APIPort string
	// APIReadTimeout is the maximum duration for reading the entire request
	APIReadTimeout time.Duration
	// APIWriteTimeout is the maximum duration before timing out writes of the response
	APIWriteTimeout time.Duration
	// APIDefaultWriteTimeout is the default timeout for write operations
	APIDefaultWriteTimeout time.Duration
	// APIIdleTimeout is the maximum amount of time to wait for the next request
	APIIdleTimeout time.Duration
	// RequestMaxBytes is the maximum size of incoming request bodies
	RequestMaxBytes int64
}

// LoadAPIConfig loads and validates API configuration from environment variables.
// It applies default values when environment variables are not set.
func LoadAPIConfig() *APIConfig {
	// Set default values.
	cnf := &APIConfig{
		BaseURL:                "/api/v0.1",
		Environment:            EnvironmentProd,
		APIHost:                "",
		APIPort:                "8080",
		APIReadTimeout:         15 * time.Second,
		APIWriteTimeout:        60 * time.Second,
		APIDefaultWriteTimeout: 30 * time.Second,
		APIIdleTimeout:         60 * time.Second,
		// common values default is 8k, Other defaults might be 4k, 16k, or 48k.
		RequestMaxBytes: 8192,
	}

	// Read and validate environment variables.
	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		cnf.Environment = env
	}

	cnf.APIHost = os.Getenv("API_HOST")

	var err error
	var parsedDuration time.Duration
	portStr := os.Getenv("API_PORT")
	_, err = strconv.Atoi(portStr) // Validate that port is #
	if err == nil {
		cnf.APIPort = portStr
	} // else if there is an error this will use the default value.

	readTimeoutStr := os.Getenv("API_READ_TIMEOUT")
	if readTimeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(readTimeoutStr); err == nil {
			cnf.APIReadTimeout = parsedDuration
		}
	}

	writeTimeoutStr := os.Getenv("API_WRITE_TIMEOUT")
	if writeTimeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(writeTimeoutStr); err == nil {
			cnf.APIWriteTimeout = parsedDuration
		}
	}

	timeoutStr := os.Getenv("API_DEFAULT_WRITE_TIMEOUT")
	if timeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(timeoutStr); err == nil {
			cnf.APIDefaultWriteTimeout = parsedDuration
		}
	}

	if cnf.APIDefaultWriteTimeout > cnf.APIWriteTimeout {
		cnf.APIDefaultWriteTimeout = cnf.APIWriteTimeout
	}

	idleTimeoutStr := os.Getenv("API_IDLE_TIMEOUT")
	if idleTimeoutStr != "" {
		if parsedDuration, err = time.ParseDuration(idleTimeoutStr); err == nil {
			cnf.APIIdleTimeout = parsedDuration
		}
	}

	maxBytesStr := os.Getenv("API_REQUEST_MAX_BYTES")
	maxBytes, err := strconv.ParseInt(maxBytesStr, 10, 64)
	if err == nil && maxBytes > 0 {
		cnf.RequestMaxBytes = maxBytes
	} // else if there is an error this will use the default value.

	return cnf
}

// LoadDBConfig loads and validates database configuration from environment variables.
// It returns a pgxpool.Config with connection parameters and pool settings.
func LoadDBConfig() *pgxpool.Config {
	pCfg, _ := pgxpool.ParseConfig("")
	pCfg.ConnConfig.Host = os.Getenv("DATASTORE_HOST")

	pCfg.ConnConfig.Port = 5432
	portStr := os.Getenv("DATASTORE_PORT")
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			pCfg.ConnConfig.Port = uint16(port)
		}
	} // else if there is an error this will use the default value.

	pCfg.ConnConfig.Database = "postgres" //nolint:goconst // These just happen to all be the same value.
	database := os.Getenv("POSTGRESQL_DATABASE")
	if database != "" {
		pCfg.ConnConfig.Database = database
	}

	pCfg.ConnConfig.User = "postgres"
	username := os.Getenv("POSTGRESQL_USERNAME")
	if username != "" {
		pCfg.ConnConfig.User = username
	}

	pCfg.ConnConfig.Password = "postgres"
	password := os.Getenv("POSTGRESQL_PASSWORD")
	if password != "" {
		pCfg.ConnConfig.Password = password
	}

	// https://github.com/jackc/pgx/blob/v5.7.2/pgxpool/pool.go
	// pgxpool.ParseConfig already sets defaults. We only need to overwrite if we set in environment variables.

	// MaxConns is the maximum size of the pool.
	// integer greater than 0 (default is the greater of 4 or runtime.NumCPU())
	maxConnsStr := os.Getenv("POOL_MAX_CONNS")
	if maxConnsStr != "" {
		maxConns, err := strconv.Atoi(maxConnsStr) // Validate that maxConns is #
		if err == nil && maxConns > 0 {
			pCfg.MaxConns = int32(maxConns)
		}
	}

	// MinConns is the minimum size of the pool. After connection closes, the pool might dip below MinConns. A low
	// number of MinConns might mean the pool is empty after MaxConnLifetime until the health check has a chance
	// to create new connections.
	// integer 0 or greater (default 0)
	minConnsStr := os.Getenv("POOL_MIN_CONNS")
	if minConnsStr != "" {
		minConns, err := strconv.Atoi(minConnsStr) // Validate that minConns is #
		if err == nil && minConns >= 0 {
			pCfg.MinConns = int32(minConns)
		}
	}

	// MaxConnIdleTime is the duration after which an idle connection will be automatically closed by the health check.
	// duration string (default 30 minutes)
	maxConnIdleTimeStr := os.Getenv("POOL_MAX_CONN_IDLE_TIME")
	if maxConnIdleTimeStr != "" {
		maxConnIdleTime, err := time.ParseDuration(maxConnIdleTimeStr)
		if err == nil {
			pCfg.MaxConnIdleTime = maxConnIdleTime
		}
	}

	// MaxConnLifetime is the duration since creation after which a connection will be automatically closed.
	// duration string (default 1 hour)
	maxConnLifetimeStr := os.Getenv("POOL_MAX_CONN_LIFETIME")
	if maxConnLifetimeStr != "" {
		maxConnLifetime, err := time.ParseDuration(maxConnLifetimeStr)
		if err == nil {
			pCfg.MaxConnLifetime = maxConnLifetime
		}
	}

	// MaxConnLifetimeJitter is the duration after MaxConnLifetime to randomly decide to close a connection.
	// This helps prevent all connections from being closed at the exact same time, starving the pool.
	// duration string (default 0)
	maxConnLifetimeJitterStr := os.Getenv("POOL_MAX_CONN_LIFETIME_JITTER")
	if maxConnLifetimeJitterStr != "" {
		maxConnLifetimeJitter, err := time.ParseDuration(maxConnLifetimeJitterStr)
		if err == nil {
			pCfg.MaxConnLifetimeJitter = maxConnLifetimeJitter
		}
	}

	// HealthCheckPeriod is the duration between checks of the health of idle connections.
	// duration string (default 1 minute)
	healthCheckPeriodStr := os.Getenv("POOL_HEALTH_CHECK_PERIOD")
	if healthCheckPeriodStr != "" {
		healthCheckPeriod, err := time.ParseDuration(healthCheckPeriodStr)
		if err == nil {
			pCfg.HealthCheckPeriod = healthCheckPeriod
		}
	}

	return pCfg
}
