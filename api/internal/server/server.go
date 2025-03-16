package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"api/internal/database"
)

// NewServer creates a new HTTP server
// [The NewServer constructor](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-newserver-constructor)
func NewServer(logger *log.Logger, db *database.Postgres) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger, db)

	// how to handle route groups?
	// how to apply middleware to some routes but not others?
	var handler http.Handler = mux
	handler = loggingMiddleware(logger, handler)
	handler = corsMiddleware(logger, handler)

	return handler
}

// encode encodes the response as JSON
// [Handle decoding/encoding in one place](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#handle-decodingencoding-in-one-place)
// TODO why is encode() and decode() different than article as written
func encode(w http.ResponseWriter, r *http.Request, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// decode decodes the request body
func decode(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("reading body: %w", err)
	}

	// Log the received body
	log.Printf("Received body: %s", string(body))

	// Check for common JSON formatting issues
	if len(body) == 0 {
		return fmt.Errorf("empty request body")
	}

	if body[0] == '\'' {
		return fmt.Errorf("invalid JSON: use double quotes (\") instead of single quotes (')")
	}

	if !strings.Contains(string(body), "\"") {
		return fmt.Errorf("invalid JSON: property names and string values must be enclosed in double quotes")
	}

	// Try to decode
	err = json.NewDecoder(bytes.NewReader(body)).Decode(v)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "looking for beginning of value"):
			return fmt.Errorf("invalid JSON format: please check your JSON syntax")
		case strings.Contains(err.Error(), "cannot unmarshal"):
			return fmt.Errorf("invalid data type in JSON: %v", err)
		default:
			return fmt.Errorf("JSON decode error: %v", err)
		}
	}

	return nil
}
