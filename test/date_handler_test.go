package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"groupie-tracker/internal/api"
	"groupie-tracker/internal/handlers"
)

// var data api.Data // Ensure this variable is declared at the package level

func TestDateHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dates", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Set a mock template for testing
	handlers.SetMockDateTemplate("<h1>{{range .Index}}<div>{{.BandName}}</div>{{end}}</h1>")
	handlers.SetMockDateTemplateError(false)

	// Populate mock data
	handlers.Data = api.Data{
		Dates: api.DateData{
			Index: []api.Date{
				{Id: 1},
				{Id: 2},
			},
		},
		Artists: []api.Artist{
			{Id: 1, Name: "Band One"},
			{Id: 2, Name: "Band Two"},
		},
	}

	handler := http.HandlerFunc(handlers.DateHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "<div>Band One</div><div>Band Two</div>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
