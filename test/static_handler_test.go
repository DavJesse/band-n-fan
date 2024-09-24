package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie-tracker/internal/handlers"
)

func TestStaticHandlerNonGetMethod(t *testing.T) {
	// Create new http request
	req, err := http.NewRequest("GET", "/static/test.css", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Make handlers.StaticHandlers act as a real HTTP server
	rr := httptest.NewRecorder()                        // Create variable to record response
	handler := http.HandlerFunc(handlers.StaticHandler) // Convert handlers.StaticHandler to http.Handler

	handler.ServeHTTP(rr, req) // Start server

	// Check if server yields bad request status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "Bad Request!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
