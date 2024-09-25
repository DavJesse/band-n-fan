package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie-tracker/internal/api"
)

var (
	artistTemplate          *template.Template
	mockArtistTemplate      string
	mockArtistTemplateError bool
)

func SetMockArtistTemplate(content string) {
	mockArtistTemplate = content
}

func SetMockArtistTemplateError(shouldError bool) {
	mockArtistTemplateError = shouldError
}

func loadArtistTemplate() error {
	if mockArtistTemplateError {
		return errors.New("mock artist template loading error")
	}
	if mockArtistTemplate != "" {
		var err error
		artistTemplate, err = template.New("artist").Parse(mockArtistTemplate)
		return err
	}
	var err error
	artistTemplate, err = template.ParseFiles("web/templates/artist.html")
	return err
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		BadRequestHandler(w)
		log.Println("Bad client request: not GET")
		return
	}

	if err := loadArtistTemplate(); err != nil {
		InternalServerErrorHandler(w)
		log.Println("Failed to load artist template", err)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		BadRequestHandler(w)
		log.Println("Error finding Artist ID")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		BadRequestHandler(w)
		log.Println("Invalid artist ID request")
		return
	}

	var selectedArtist api.Artist
	var foundArtist, foundRelation bool

	// Extract artist with matching id
	for _, artist := range Data.Artists {
		if artist.Id == id {
			selectedArtist = artist
			foundArtist = true
			break
		}
	}

	// Extract relation data for tour dates and locations
	for _, relation := range Data.Relations.Index {
		if relation.Id == id {
			selectedArtist.Relation = relation.DatesLocations
			foundRelation = true
			break
		}
	}

	if !foundArtist {
		NotFoundHandler(w)
		log.Println("Artist index not found")
		return
	} else if !foundRelation {
		NotFoundHandler(w)
		log.Println("Relation index not found")
	}

	err = artistTemplate.Execute(w, selectedArtist)
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Failed to load selected artist template", err)
		return
	}
}
