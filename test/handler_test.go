package test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
)

func TestHomeHandler(t *testing.T) {
	tempDir := t.TempDir()

	tmplFile := tempDir + "/home.html"
	err := os.WriteFile(tmplFile, []byte("<html><body>Home Page</body></html>"), 0644)

	if err != nil {
		t.Fatalf("Error creating template files: %v", err)
	}

	origParseFiles := template.ParseFiles
	template.ParseFiles = func(filenames ...string) (*template.Template, error) {
		return origParseFiles(tmplFile)
	}
	defer func() {
		template.ParseFiles = origParseFiles
	}()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf(err)
	}
	rr := httptest.NewRecorder()
	handler := http.Handler(TestHomeHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body contains the correct content.
	expected := "<html><body>Home Page</body></html>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Test case: Non-GET request (e.g., POST)
	req, err = http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for non-GET request: got %v want %v", status, http.StatusBadRequest)
	}
}
