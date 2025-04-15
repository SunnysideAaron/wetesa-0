//nolint:noctx // context isn't needed for tests
package testintegration

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestHealthzEndpoint(t *testing.T) {
	// Get server address from config
	serverAddr := getServerAddress()

	// Create an HTTP client
	client := &http.Client{}

	// Make the request
	resp, err := client.Get(serverAddr + "/healthz")
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			t.Errorf("failed to close response body: %v", err)
		}
	}()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check response body
	if string(body) != "OK" {
		t.Errorf("expected body 'OK', got '%s'", string(body))
	}
}

func TestHealthDBzEndpoint(t *testing.T) {
	serverAddr := getServerAddress()
	client := &http.Client{}

	resp, err := client.Get(serverAddr + "/healthdbz")
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			t.Errorf("failed to close response body: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check if response contains "status" field
	if !strings.Contains(string(body), "status") {
		t.Errorf("response body does not contain 'status' field: %s", string(body))
	}
}
