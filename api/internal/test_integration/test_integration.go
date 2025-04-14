package test_integration

import (
	"fmt"

	"api/internal/config"
)

func getServerAddress() string {
	cfg := config.LoadAPIConfig()
	host := cfg.APIHost
	if host == "" {
		host = "localhost"
	}
	return fmt.Sprintf("http://%s:%s", host, cfg.APIPort)
}

func getAPIAddress() string {
	cfg := config.LoadAPIConfig()
	host := cfg.APIHost
	if host == "" {
		host = "localhost"
	}
	return fmt.Sprintf("http://%s:%s%s", host, cfg.APIPort, cfg.BaseURL)
}
