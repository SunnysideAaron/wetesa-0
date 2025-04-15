// Package testintegration provides integration tests for the API.
// Note that server has to be running for these tests to pass.
package testintegration

import (
	"net"

	"api/internal/config"
)

func getServerAddress() string {
	cfg := config.LoadAPIConfig()
	host := cfg.APIHost
	if host == "" {
		host = "localhost"
	}
	return "http://" + net.JoinHostPort(host, cfg.APIPort)
}

func getAPIAddress() string {
	cfg := config.LoadAPIConfig()
	host := cfg.APIHost
	if host == "" {
		host = "localhost"
	}
	return "http://" + net.JoinHostPort(host, cfg.APIPort) + cfg.BaseURL
}
