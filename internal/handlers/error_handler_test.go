package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBadRequestHandler(t *testing.T) {
	// Override loadTemplate for test purposes
	loadTemplate = func() (*template.Template, error) {
		return template.New("test").Parse("{{.Problem}}")
	}

	rec := httptest.NewRecorder()

	// Call the handler
	badRequestHandler(rec)

	// Check the status code
	if status := rec.Code; status != http.StatusOK{
		t.Errorf("expected status code %d, but got %d", http.StatusBadRequest, status)
	}

	// Check the response body
	expected := "Bad Request!"
	if !strings.Contains(rec.Body.String(), expected) {
		t.Errorf("expected response body to contain %q, but got %q", expected, rec.Body.String())
	}
}

func TestInternalServerErrorHandler(t *testing.T) {
	// Override loadTemplate for test purposes
	loadTemplate = func() (*template.Template, error) {
		return template.New("test").Parse("{{.Problem}}")
	}

	rec := httptest.NewRecorder()

	// Call the handler
	internalServerErrorHandler(rec)

	// Check the status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusInternalServerError, status)
	}

	// Check the response body
	expected := "Internal Server Error!"
	if !strings.Contains(rec.Body.String(), expected) {
		t.Errorf("expected response body to contain %q, but got %q", expected, rec.Body.String())
	}
}

func TestNotFoundHandler(t *testing.T) {
	// Override loadTemplate for test purposes
	loadTemplate = func() (*template.Template, error) {
		return template.New("test").Parse("{{.Problem}}")
	}

	rec := httptest.NewRecorder()

	// Call the handler
	notFoundHandler(rec)

	// Check the status code
	if status := rec.Code; status != http.StatusOK{
		t.Errorf("expected status code %d, but got %d", http.StatusNotFound, status)
	}

	// Check the response body
	expected := "Not Found!"
	if !strings.Contains(rec.Body.String(), expected) {
		t.Errorf("expected response body to contain %q, but got %q", expected, rec.Body.String())
	}
}

