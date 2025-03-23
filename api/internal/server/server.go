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

	// TODO most "middlewares" need to go before routes. and order matters for example logging before auth.
	// TODO put middlewates into addRoutes
	// TODO how to handle route groups?
	// TODO how to apply middleware to some routes but not others?
	//var handler http.Handler = mux
	// handler = loggingMiddleware(logger, handler)
	// handler = corsMiddleware(logger, handler)

	return mux
}

// encode encodes the response as JSON
// [Handle decoding/encoding in one place](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#handle-decodingencoding-in-one-place)
func encode(w http.ResponseWriter, r *http.Request, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		//TODO logging / error handling
		// we can probably refactor this to do the http error return. here
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

// [Validating data](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#validating-data)
func decode[T database.Validator](r *http.Request) (T, map[string]string, error) {
	var v T

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return v, nil, fmt.Errorf("reading body: %w", err)
	}

	// Log the received body
	log.Printf("Received body: %s", string(body))

	// Check for common JSON formatting issues
	if len(body) == 0 {
		return v, nil, fmt.Errorf("empty request body")
	}

	if body[0] == '\'' {
		return v, nil, fmt.Errorf("invalid JSON: use double quotes (\") instead of single quotes (')")
	}

	if !strings.Contains(string(body), "\"") {
		return v, nil, fmt.Errorf("invalid JSON: property names and string values must be enclosed in double quotes")
	}

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&v)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "looking for beginning of value"):
			return v, nil, fmt.Errorf("invalid JSON format: please check your JSON syntax")
		case strings.Contains(err.Error(), "cannot unmarshal"):
			return v, nil, fmt.Errorf("invalid data type in JSON: %v", err)
		default:
			return v, nil, fmt.Errorf("JSON decode error: %v", err)
		}
	}

	if problems := v.Valid(r.Context()); len(problems) > 0 {
		return v, problems, fmt.Errorf("invalid %T: %d problems", v, len(problems))
	}

	return v, nil, nil
}
