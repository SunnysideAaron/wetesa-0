//nolint:noctx // context isn't needed for tests
package testintegration

import (
	"io"
	"net/http"
	"testing"
)

func TestLogLevelEndpoint(t *testing.T) {
	serverAddr := getAPIAddress()
	client := &http.Client{}

	testCases := []struct {
		name           string
		level          string
		expectedStatus int
		expectedBody   string
	}{
		{"Debug Level", "debug", http.StatusOK, "OK"},
		{"Info Level", "info", http.StatusOK, "OK"},
		{"Invalid Level", "invalid", http.StatusOK, "OK"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.Get(serverAddr + "/loglevel/" + tc.level)
			if err != nil {
				t.Fatalf("failed to make request: %v", err)
			}
			defer func() {
				if err = resp.Body.Close(); err != nil {
					t.Errorf("failed to close response body: %v", err)
				}
			}()

			if resp.StatusCode != tc.expectedStatus {
				t.Errorf("expected status code %d, got %d", tc.expectedStatus, resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}

			if string(body) != tc.expectedBody {
				t.Errorf("expected body '%s', got: '%s'", tc.expectedBody, string(body))
			}
		})
	}
}
