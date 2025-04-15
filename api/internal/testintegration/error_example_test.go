//nolint:noctx // context isn't needed for tests
package testintegration

import (
	"io"
	"net/http"
	"testing"
)

func TestErrorExampleEndpoint(t *testing.T) {
	serverAddr := getAPIAddress()
	client := &http.Client{}

	resp, err := client.Get(serverAddr + "/errorexample")
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			t.Errorf("failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedMsg := "i'm a demon sent to torment you 😈\n"
	if string(body) != expectedMsg {
		t.Errorf("expected error message %q, got: %q", expectedMsg, string(body))
	}
}
