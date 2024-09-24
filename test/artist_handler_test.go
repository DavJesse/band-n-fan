package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"groupie-tracker/internal/handlers"
)

func TestArtistHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/?id=1", nil) // Example with id=1
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Set a mock template for testing
	handlers.SetMockArtistTemplate("<h1>{{.Name}}</h1>") // Example mock template
	handlers.SetMockArtistTemplateError(false)           // Ensure no mock error is set

	handler := http.HandlerFunc(handlers.ArtistHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "<h1>Queen</h1>" // Replace with actual expected output based on the mock
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
