package handlers

import (
	"encoding/json"
	"groupie-tracker/internal/api"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	ids := SearchArtist(query)       // Retrieve results
	results := GetResults(ids, Data) //Retrieve results

	// Execute tmpl with search query
	err = tmpl.Execute(w, results)
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Failed to execute search template:", err)
		return
	}
}

func SearchArtist(query string) []int {
	var resultIDs []int
	query = strings.ToLower(query)

	// Search for matching artist
	for _, artist := range Data.Artists {
		// Search by name
		if strings.Contains(strings.ToLower(artist.Name), query) && !IdExists(resultIDs, artist.Id) {
			resultIDs = append(resultIDs, artist.Id)
		}
		// Search by FirstAlbum Date
		if strings.Contains(artist.FirstAlbum, query) && !IdExists(resultIDs, artist.Id) {
			resultIDs = append(resultIDs, artist.Id)
		}
		// Search by creation date
		if IsNumeric(query) {
			date, _ := strconv.Atoi(query)
			if date == artist.CreationDate && !IdExists(resultIDs, artist.Id) {
				resultIDs = append(resultIDs, artist.Id)
			}
		}

		// Search by band members
		for i := range artist.Members {
			if strings.Contains(strings.ToLower(artist.Members[i]), query) && !IdExists(resultIDs, artist.Id) {
				resultIDs = append(resultIDs, artist.Id)
				break
			}
		}
	}

	// Search for matching date
	for _, dateObj := range Data.Dates.Index {
		for _, date := range dateObj.Dates {
			if strings.Contains(date, query) && !IdExists(resultIDs, dateObj.Id) {
				resultIDs = append(resultIDs, dateObj.Id)
			}
		}
	}

	// Search for matching locations
	for _, locationObj := range Data.Locations.Index {
		for _, location := range locationObj.Locations {
			if strings.Contains(location, query) && !IdExists(resultIDs, locationObj.Id) {
				resultIDs = append(resultIDs, locationObj.Id)
			}
		}
	}

	return resultIDs
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

// Check if id exists in a slice of ids
func IdExists(ids []int, id int) bool {
	for i := range ids {
		if ids[i] == id {
			return true
		}
	}
	return false
}

func GetResults(ids []int, data api.Data) []api.Artist {
	var results []api.Artist

	for i := range ids {
		for _, artist := range data.Artists {
			if ids[i] == artist.Id {
				results = append(results, artist)
				break
			}
		}
	}
	return results
}

func SuggestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q") // Retrieve search query from html form
	results := SearchArtist(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err = template.ParseFiles("web/templates/search_results.html")
// 	if err != nil {
// 		InternalServerErrorHandler(w)
// 		log.Println("Failed to load search template:", err)
// 		return
// 	}
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		InternalServerErrorHandler(w)
// 		log.Println("Failed to load search template:", err)
// 		return
// 	}
// }
