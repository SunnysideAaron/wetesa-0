package config

import (
	"os"
	"strconv"
	"time"
)

// APIConfig stores the api configuration.
type APIConfig struct {
	APIHost         string
	APIPort         string
	APIReadTimeout  time.Duration
	APIWriteTimeout time.Duration
	APIIdleTimeout  time.Duration
}

// LoadAPIConfig loads the api configuration from environment variables. Default values are
// used if not set or invalid. Values are validated if appropriate.
func LoadAPIConfig() *APIConfig {
	// Set default values.
	cnf := &APIConfig{
		APIPort:         "8080",
		APIReadTimeout:  15,
		APIWriteTimeout: 30,
		APIIdleTimeout:  60,
	}

	// Read and validate environment variables.
	cnf.APIHost = os.Getenv("API_HOST")

	portStr := os.Getenv("API_PORT")
	_, err := strconv.Atoi(portStr) //Validate that port is #
	if err == nil {
		cnf.APIPort = portStr
	} // else if there is an error this will use the default value.

	readTimeoutStr := os.Getenv("API_PORT")
	readTimeout, err := strconv.Atoi(readTimeoutStr) //Validate that readTimeoutStr is #
	if err == nil {
		cnf.APIReadTimeout = time.Duration(readTimeout)
	} // else if there is an error this will use the default value.

	writeTimeoutStr := os.Getenv("API_PORT")
	writeTimeout, err := strconv.Atoi(writeTimeoutStr) //Validate that writeTimeoutStr is #
	if err == nil {
		cnf.APIWriteTimeout = time.Duration(writeTimeout)
	} // else if there is an error this will use the default value.

	idleTimeoutStr := os.Getenv("API_PORT")
	idleTimeout, err := strconv.Atoi(idleTimeoutStr) //Validate that idleTimeoutStr is #
	if err == nil {
		cnf.APIIdleTimeout = time.Duration(idleTimeout)
	} // else if there is an error this will use the default value.

	// TODO log values if set to debug
	// log.Printf("APIHost: %s", cnf.APIHost)
	// log.Printf("APIPort: %s", cnf.APIPort)

	return cnf
}

// // DBConfig stores the database configuration.
// type DBConfig struct {
// 	APIHost         string
// 	APIPort         string
// 	APIReadTimeout  time.Duration
// 	APIWriteTimeout time.Duration
// 	APIIdleTimeout  time.Duration
// }

// // LoadDBConfig loads the db configuration from environment variables. Default values are
// // used if not set or invalid. Values are validated if appropriate.
// func LoadDBConfig() *APIConfig {
// 	// Set default values.
// 	cnf := &DBConfig{
// 		APIPort:         "8080",
// 		APIReadTimeout:  15,
// 		APIWriteTimeout: 30,
// 		APIIdleTimeout:  60,
// 	}

// 	// Read and validate environment variables.
// 	cnf.APIHost = os.Getenv("API_HOST")

// 	portStr := os.Getenv("API_PORT")
// 	_, err := strconv.Atoi(portStr) //Validate that port is #
// 	if err == nil {
// 		cnf.APIPort = portStr
// 	} // else if there is an error this will use the default value.

// 	readTimeoutStr := os.Getenv("API_PORT")
// 	readTimeout, err := strconv.Atoi(readTimeoutStr) //Validate that readTimeoutStr is #
// 	if err == nil {
// 		cnf.APIReadTimeout = time.Duration(readTimeout)
// 	} // else if there is an error this will use the default value.

// 	writeTimeoutStr := os.Getenv("API_PORT")
// 	writeTimeout, err := strconv.Atoi(writeTimeoutStr) //Validate that writeTimeoutStr is #
// 	if err == nil {
// 		cnf.APIWriteTimeout = time.Duration(writeTimeout)
// 	} // else if there is an error this will use the default value.

// 	idleTimeoutStr := os.Getenv("API_PORT")
// 	idleTimeout, err := strconv.Atoi(idleTimeoutStr) //Validate that idleTimeoutStr is #
// 	if err == nil {
// 		cnf.APIIdleTimeout = time.Duration(idleTimeout)
// 	} // else if there is an error this will use the default value.

// 	// TODO log values if set to debug
// 	// log.Printf("APIHost: %s", cnf.APIHost)
// 	// log.Printf("APIPort: %s", cnf.APIPort)

// 	return cnf
// }
