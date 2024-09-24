package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request for path: %s\n", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/artists":
			json.NewEncoder(w).Encode([]Artist{
				{Id: 1, Name: "The Beatles", Members: []string{"John Lennon", "Paul McCartney"}},
			})
		case "/api/locations":
			json.NewEncoder(w).Encode(LocationData{Index: []Location{{Id: 1, Locations: []string{"Liverpool"}, BandName: "The Beatles"}}})
		case "/api/dates":
			json.NewEncoder(w).Encode(DateData{Index: []Date{{Id: 1, Dates: []string{"1960-08-01"}, BandName: "The Beatles"}}})
		case "/api/relation":
			json.NewEncoder(w).Encode(RelationData{Index: []Relation{{Id: 1, DatesLocations: map[string][]string{"Liverpool": {"1960-08-01"}}, BandName: "The Beatles"}}})
		default:
			http.NotFound(w, r)
		}
	}))
}

func TestFetchData(t *testing.T) {
	mockServer := setupMockServer()
	defer mockServer.Close()

	var artists []Artist
	url := mockServer.URL + "/api/artists"
	fmt.Printf("Fetching data from: %s\n", url)
	err := fetchData(url, &artists)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(artists) != 1 || artists[0].Name != "The Beatles" {
		t.Errorf("expected artist name to be 'The Beatles', got %v", artists[0].Name)
	}
}

func TestLoadData(t *testing.T) {
	mockServer := setupMockServer()
	defer mockServer.Close()

	// Save the original apiURL and restore it after the test
	originalAPIURL := apiURL
	apiURL = mockServer.URL + "/api"
	defer func() { apiURL = originalAPIURL }()

	fmt.Printf("Using API URL: %s\n", apiURL)

	data, err := LoadData()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test Artists
	if len(data.Artists) != 1 || data.Artists[0].Name != "The Beatles" {
		t.Errorf("expected artist name to be 'The Beatles', got %v", data.Artists[0].Name)
	}

	// Test Locations
	if len(data.Locations.Index) != 1 || data.Locations.Index[0].BandName != "The Beatles" {
		t.Errorf("expected location band name to be 'The Beatles', got %v", data.Locations.Index[0].BandName)
	}

	// Test Dates
	if len(data.Dates.Index) != 1 || data.Dates.Index[0].BandName != "The Beatles" {
		t.Errorf("expected date band name to be 'The Beatles', got %v", data.Dates.Index[0].BandName)
	}

	// Test Relations
	if len(data.Relations.Index) != 1 || data.Relations.Index[0].BandName != "The Beatles" {
		t.Errorf("expected relation band name to be 'The Beatles', got %v", data.Relations.Index[0].BandName)
	}
}

func TestFetchDataError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer mockServer.Close()

	var artists []Artist
	url := mockServer.URL + "/api/artists"
	fmt.Printf("Fetching data from (error test): %s\n", url)
	err := fetchData(url, &artists)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
