package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/api"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict acces to '/search' page
	if r.Method != "GET" {
		BadRequestHandler(w)
		log.Println("Search handler bad error")
		return
	}

	// Update tmpl with http template of search.html
	tmpl, err = template.ParseFiles("web/templates/search_results.html")
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Failed to load search template:", err)
		return
	}

	query := r.URL.Query().Get("artist") // Retrieve search query from html form

	results := SearchArtist(query) // Retrieve results

	// Execute tmpl with search query
	err = tmpl.Execute(w, results)
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Failed to execute search template:", err)
		return
	}
}

func SearchArtist(query string) api.Data {
	var result api.Data
	query = strings.ToLower(query)

	// Search for matching artist
	for _, artist := range Data.Artists {
		// Search by name
		if strings.Contains(artist.Name, query) {
			result.Artists = append(result.Artists, artist)
		}
		// Search by FirstAlbum Date
		if strings.Contains(artist.FirstAlbum, query) {
			result.Artists = append(result.Artists, artist)
		}
		// Search by creation date
		if IsNumeric(query) {
			date, _ := strconv.Atoi(query)
			if date == artist.CreationDate {
				result.Artists = append(result.Artists, artist)
			}
		}

		// Search by band members
		for i := range artist.Members {
			if strings.Contains(artist.Members[i], query) {
				result.Artists = append(result.Artists, artist)
				break
			}
		}
	}

	// Search for matching date
	for _, dateObj := range Data.Dates.Index {
		for _, date := range dateObj.Dates {
			if strings.Contains(date, query) {
				result.Dates.Index = append(result.Dates.Index, dateObj)
			}
		}
	}

	// Search for matching locations
	for _, locationObj := range Data.Locations.Index {
		for _, location := range locationObj.Locations {
			if strings.Contains(location, query) {
				result.Locations.Index = append(result.Locations.Index, locationObj)
			}
		}
	}

	return result
}

func IsNumeric(str string) bool {
	// Search for instances of non-numeric characters
	for _, ch := range str {
		if !(ch >= '0' && ch <= '9') {
			return false
		}
	}
	return true
}

func SuggestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.Url.Query().Get("q") // Retrieve search query from html form
	results := SearchArtist(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err = template.ParseFiles("web/templates/search_results.html")
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Failed to load search template:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Failed to load search template:", err)
		return
	}
}
