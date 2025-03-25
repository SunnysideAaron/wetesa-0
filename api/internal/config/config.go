// [The Twelve-Factor App III. Config](https://12factor.net/config)
package config

import (
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	EnvironmentDev  string = "dev"
	EnvironmentProd string = "prod"
)

// APIConfig stores the api configuration.
type APIConfig struct {
	Environment     string
	APIHost         string
	APIPort         string
	APIReadTimeout  time.Duration
	APIWriteTimeout time.Duration
	APIIdleTimeout  time.Duration
	RequestMaxBytes int64
	RequestTimeout  time.Duration
}

// LoadAPIConfig loads the api configuration from environment variables. Default values are
// used if not set or invalid. Values are validated if appropriate.
func LoadAPIConfig() *APIConfig {
	// Set default values.
	cnf := &APIConfig{
		Environment:     EnvironmentProd,
		APIPort:         "8080",
		APIReadTimeout:  15,
		APIWriteTimeout: 30,
		APIIdleTimeout:  60,
		// common values default is 8k, Other defaults might be 4k, 16k, or 48k.
		RequestMaxBytes: 8192,
		RequestTimeout:  60,
	}

	// Read and validate environment variables.
	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		cnf.Environment = env
	}

	cnf.APIHost = os.Getenv("API_HOST")

	portStr := os.Getenv("API_PORT")
	_, err := strconv.Atoi(portStr) //Validate that port is #
	if err == nil {
		cnf.APIPort = portStr
	} // else if there is an error this will use the default value.

	readTimeoutStr := os.Getenv("API_READ_TIME_OUT")
	readTimeout, err := strconv.Atoi(readTimeoutStr) //Validate that readTimeoutStr is #
	if err == nil {
		cnf.APIReadTimeout = time.Duration(readTimeout)
	} // else if there is an error this will use the default value.

	writeTimeoutStr := os.Getenv("API_WRITE_TIME_OUT")
	writeTimeout, err := strconv.Atoi(writeTimeoutStr) //Validate that writeTimeoutStr is #
	if err == nil {
		cnf.APIWriteTimeout = time.Duration(writeTimeout)
	} // else if there is an error this will use the default value.

	idleTimeoutStr := os.Getenv("API_IDLE_TIME_OUT")
	idleTimeout, err := strconv.Atoi(idleTimeoutStr) //Validate that idleTimeoutStr is #
	if err == nil {
		cnf.APIIdleTimeout = time.Duration(idleTimeout)
	} // else if there is an error this will use the default value.

	maxBytesStr := os.Getenv("API_REQUEST_MAX_BYTES")
	maxBytes, err := strconv.ParseInt(maxBytesStr, 10, 64)
	if err == nil && maxBytes > 0 {
		cnf.RequestMaxBytes = maxBytes
	} // else if there is an error this will use the default value.

	timeoutStr := os.Getenv("API_REQUEST_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err == nil && timeout > 0 {
		cnf.RequestTimeout = time.Duration(timeout)
	}

	// TODO log values if set to debug
	// log.Printf("APIHost: %s", cnf.APIHost)
	// log.Printf("APIPort: %s", cnf.APIPort)

	return cnf
}

// LoadDBConfig loads the database configuration from environment variables.
// Default values are used if not set or invalid.
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

	pCfg.ConnConfig.Database = "postgres"
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
		maxConns, err := strconv.Atoi(maxConnsStr) //Validate that maxConns is #
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
		minConns, err := strconv.Atoi(minConnsStr) //Validate that minConns is #
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

	//log.Printf("pCfg: %s", pCfg)

	return pCfg
}
