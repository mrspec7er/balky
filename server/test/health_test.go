package test

import (
	"net/http"
	"testing"
)

func TestServerHealth(t *testing.T) {
	// Replace with the actual server URL and endpoint
	serverURL := "http://localhost:8080/"

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, serverURL, nil)
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request and handle the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if expected := http.StatusOK; resp.StatusCode != expected {
		t.Errorf("unexpected response status code: expected %d, got %d", expected, resp.StatusCode)
	}
}
