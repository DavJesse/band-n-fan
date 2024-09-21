package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie-tracker/internal/api"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/artist.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to load artist template", err)
		return
	}

	// Parse 'id' query as parameter
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		badRequestHandler(w)
		log.Println("Error finding Artist ID")
		return
	}

	// Convert 'idStr' to int for practical use
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequestHandler(w)
		log.Println("Invalid artist ID request")
		return
	}

	var selectedArtist api.Artist
	var foundArtist bool
	var foundRelation bool

	// Extract artist with matching id
	for _, artist := range data.Artists {
		if artist.Id == id {
			selectedArtist = artist
			foundArtist = true
			break
		}
	}

	// Extract relation data for tour dates and locations
	for _, relation := range data.Relations.Index {
		if relation.Id == id {
			selectedArtist.Relation = relation.DatesLocations
			foundRelation = true
			break
		}
	}

	// Handle error when client request is not found
	if !foundArtist {
		notFoundHandler(w)
		log.Println("Artist index not found")
		return
	} else if !foundRelation {
		notFoundHandler(w)
		log.Println("Relation index not found")
	}

	// Render artist to html template
	err = tmpl.Execute(w, selectedArtist)
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to load selected artist template", err)
		return
	}
}
