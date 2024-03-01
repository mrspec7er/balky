package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mrspec7er/balky/app/model"
)

type GetUsersResponse struct {
	Status   bool         `json:"status"`
	Message  string       `json:"message"`
	Data     []model.User `json:"data"`
	Metadata interface{}  `json:"metadata"`
}

func TestGetAllUser(t *testing.T) {
	// Replace with the actual server URL and endpoint
	serverURL := "http://localhost:8080/users"

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

	var expectedResponse GetUsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&expectedResponse); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	// Assert the response code
	if expected := http.StatusOK; resp.StatusCode != expected {
		t.Errorf("unexpected response status code: expected %d, got %d", expected, resp.StatusCode)
	}

}
