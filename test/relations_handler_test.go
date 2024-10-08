package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"groupie-tracker/internal/api"
	"groupie-tracker/internal/handlers"
)

func TestRelationsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/relations", nil) // Adjust the URL if needed
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Set a mock template for testing
	handlers.SetMockRelationTemplate("<h1>{{range .Index}}<div>{{.BandName}}</div>{{end}}</h1>")
	handlers.SetMockRelationTemplateError(false)

	// Populate mock data
	handlers.Data = api.Data{
		Relations: api.RelationData{
			Index: []api.Relation{
				{Id: 1},
				{Id: 2},
			},
		},
		Artists: []api.Artist{
			{Id: 1, Name: "Queen"},
			{Id: 2, Name: "Soja"},
		},
	}

	handler := http.HandlerFunc(handlers.RelationsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "<div>Queen</div><div>Soja</div>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
