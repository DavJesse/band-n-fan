package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie-tracker/internal/api"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// if r.method != post {
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// 	return
	// }

	tmpl, err := template.ParseFiles("web/templates/artist.html")
	if err != nil {
		http.Error(w, "Failed to load artist", http.StatusInternalServerError)
		log.Println("Error loading template: ", err)
		return
	}

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
		log.Println("Error finding Artist ID")
		return
	}

	// Convert 'idStr' to int for practical use
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	var selectedArtist api.Artist
	var foundArtist bool
	var foundRelation bool

	// Extract artist with matching id
	for _, artist := range data.Artists {
		if artist.Id == id {
			log.Println("Artist from data.Artists")
			log.Printf("%#v\n", artist)
			selectedArtist = artist
			foundArtist = true
			break
		}
	}

	for _, relation := range data.Relations.Index {
		if relation.Id == id {
			selectedArtist.Relation = relation.DatesLocations
			foundRelation = true
			break
		}
	}

	if !foundArtist {
		http.Error(w, "Artist not found", http.StatusNotFound)
		log.Println("Artist index not found")
		return
	} else if !foundRelation {
		http.Error(w, "Relation not found", http.StatusNotFound)
		log.Println("Relation index not found")
	}

	// Render artist to html template
	err = tmpl.Execute(w, selectedArtist)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Println("Failed to load selected artist template", err)
		return
	}
}
