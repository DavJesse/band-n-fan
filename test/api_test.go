package test

import (
	"net/http"
	"net/http/httptest"
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
