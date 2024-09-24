package test

import (
	"testing"

	"groupie-tracker/internal/api"
)

func TestLoadData(t *testing.T) {
	data, err := api.LoadData()
	// Check for errors yielded during loading
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	// Check for failure to load all artists
	if len(data.Artists) != 52 {
		t.Error("No artists loaded")
	}

	// Check for failure to load all locations
	if len(data.Locations.Index) != 52 {
		t.Error("No locations loaded")
	}

	// Check for failure to all load dates
	if len(data.Dates.Index) != 52 {
		t.Error("No dates loaded")
	}

	// Check for failure to load all relations
	if len(data.Relations.Index) != 52 {
		t.Error("No relations loaded")
	}

	// Check if the first artist has expected fields
	if data.Artists[0].Id == 0 || data.Artists[0].Name == "" {
		t.Error("Invalid artist data loaded")
	}

	// Check if the first location has expected fields
	if data.Locations.Index[0].Id == 0 || len(data.Locations.Index[0].Locations) == 0 {
		t.Error("Invalid location data loaded")
	}

	// Check if the first date has expected fields
	if data.Dates.Index[0].Id == 0 || len(data.Dates.Index[0].Dates) == 0 {
		t.Error("Invalid date data loaded")
	}

	// Check if the first relation has expected fields
	if data.Relations.Index[0].Id == 0 || len(data.Relations.Index[0].DatesLocations) == 0 {
		t.Error("Invalid relation data loaded")
	}
}
