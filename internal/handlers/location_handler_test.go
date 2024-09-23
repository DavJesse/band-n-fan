package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"groupie-tracker/internal/api"
)

func TestLocationsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/locations", nil) // Adjust the URL if needed
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Set a mock template for testing
	SetMockLocationTemplate("<h1>{{range .Index}}<div>{{.BandName}}</div>{{end}}</h1>")
	SetMockLocationTemplateError(false)

	// Populate mock data
	data = api.Data{
		Locations: api.LocationData{
			Index: []api.Location{
				{Id: 1},
				{Id: 2},
			},
		},
		Artists: []api.Artist{
			{Id: 1, Name: "Artist One"},
			{Id: 2, Name: "Artist Two"},
		},
	}

	handler := http.HandlerFunc(LocationsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "<div>Artist One</div><div>Artist Two</div>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
