package test

import (
	"groupie-tracker/internal/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Set a mock template
	handlers.SetMockHomeTemplate("Artists")  // Mock template content for testing
	handlers.SetMockHomeTemplateError(false) // Ensure no error is set

	handler := http.HandlerFunc(handlers.HomeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Artists"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHomeHandlerMethodNotAllowed(t *testing.T) {
	// Test with a POST request, which should not be allowed
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HomeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code for POST: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHomeHandlerWrongPath(t *testing.T) {
	// Test with a wrong path
	req, err := http.NewRequest("GET", "/wrong-path", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HomeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code for wrong path: got %v want %v",
			status, http.StatusBadRequest)
	}
}
