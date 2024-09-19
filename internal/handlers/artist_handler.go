package handlers

import (
	"groupie-tracker/internal/api"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.method != post {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("web\templates\artist.html")

	// Load data from API
	data, err := api.LoadData()

	if err != nil {
		http.Error(w, "Failed to load data", http.StatusInternalServerError)
		return
	}

	// Parse 'id' query as parameter
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Artist ID missing", http.StatusBadRequest)
		return
	}

	// Convert 'idStr' to int for practical use
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	var selectedArtist api.Artist
	var found bool

	// Extract artist with matching id
	for _, artist := range data.Artists {
		if artist.Id == id {
			selectedArtist = artist
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Render artist to html template
	err = tmpl.Execute(w, selectedArtist)

	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}

}
