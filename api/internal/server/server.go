// Package server provides HTTP server functionality including request handling,
// middleware, and routing for the API.
package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"api/internal/database"
)

// encode writes the response as JSON to the http.ResponseWriter with the given status code.
// It returns an error if JSON encoding fails.
// [Handle decoding/encoding in one place](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#handle-decodingencoding-in-one-place)
func encode(w http.ResponseWriter, _ *http.Request, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
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

	// Check for common JSON formatting issues
	if len(body) == 0 {
		return v, nil, errors.New("empty request body")
	}

	if body[0] == '\'' {
		return v, nil, errors.New("invalid JSON - use double quotes (\") instead of single quotes (')")
	}

	if !strings.Contains(string(body), "\"") {
		return v, nil, errors.New("invalid JSON - property names and string values must be enclosed in double quotes")
	}

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&v)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "looking for beginning of value"):
			return v, nil, fmt.Errorf("invalid JSON format - please check your JSON syntax: %w", err)
		case strings.Contains(err.Error(), "cannot unmarshal"):
			return v, nil, fmt.Errorf("invalid data type in JSON: %w", err)
		default:
			return v, nil, fmt.Errorf("JSON decode error: %w", err)
		}
	}

	if problems := v.Valid(r.Context()); len(problems) > 0 {
		return v, problems, fmt.Errorf("invalid %T %d problems", v, len(problems))
	}

	return v, nil, nil
}
