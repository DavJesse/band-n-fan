package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"groupie-tracker/internal/api"
)

func TestFetchData(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "Test Artist", "id": 1}`))
	}))
	defer server.Close()

	// Create a target struct to unmarshal the data into
	var target struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}

	// Call FetchData with the mock server URL
	err := api.FetchData(server.URL, &target)
	// Check if there's no error
	if err != nil {
		t.Errorf("FetchData returned an error: %v", err)
	}

	// Check if the data was correctly unmarshaled
	expectedName := "Test Artist"
	expectedID := 1

	if target.Name != expectedName {
		t.Errorf("Expected name %s, but got %s", expectedName, target.Name)
	}

	if target.ID != expectedID {
		t.Errorf("Expected ID %d, but got %d", expectedID, target.ID)
	}
}

func TestFetchDataInvalidURL(t *testing.T) {
	var target interface{}
	invalidUrl := "http://invalid_url.com" // Setup variable an invalid URL

	err := api.FetchData(invalidUrl, &target) // Try to fetch data

	// Check if funcction yields error
	if err == nil {
		t.Error("Expected an error but got nil")
	}

	// Check if function yields the correct error message
	if !strings.Contains(err.Error(), "failed to fetch data") {
		t.Error("Expecte the error message: 'failed to fetch data from'")
		t.Error("Got: %v", err)
	}
}
