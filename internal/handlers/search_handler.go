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

type ResultIDs struct {
	QueryResult string
	SearchParam string
	Id          int
}

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
	results := GetResults(ids, Data) // Retrieve results

	// Execute tmpl with search query
	err = tmpl.Execute(w, results)
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Failed to execute search template:", err)
		return
	}
}

func SearchArtist(query string) []ResultIDs {
	var results []ResultIDs
	query = strings.ToLower(query)

	// Search for matching artist
	for _, artist := range Data.Artists {
		// Search by name
		if strings.Contains(strings.ToLower(artist.Name), query) {
			band := ResultIDs{}
			band.SearchParam = "Band Name"
			band.Id = artist.Id
			band.QueryResult = artist.Name
			results = append(results, band)
		}
		// Search by FirstAlbum Date
		if strings.Contains(artist.FirstAlbum, query) {
			band := ResultIDs{}
			band.SearchParam = "First Album Date"
			band.Id = artist.Id
			band.QueryResult = artist.FirstAlbum
			results = append(results, band)
		}
		// Search by creation date
		if IsNumeric(query) {
			date, _ := strconv.Atoi(query)
			if date == artist.CreationDate {
				band := ResultIDs{}
				band.SearchParam = "Creation Date"
				band.Id = artist.Id
				band.QueryResult = strconv.Itoa(artist.CreationDate)
				results = append(results, band)
			}
		}

		// Search by band members
		for i := range artist.Members {
			if strings.Contains(strings.ToLower(artist.Members[i]), query) {
				band := ResultIDs{}
				band.SearchParam = "Band Member"
				band.Id = artist.Id
				band.QueryResult = artist.Members[i]
				results = append(results, band)
				break
			}
		}
	}

	// Search for matching date
	for _, dateObj := range Data.Dates.Index {
		for _, date := range dateObj.Dates {
			if strings.Contains(date, query) {
				band := ResultIDs{}
				band.SearchParam = "Tour Date"
				band.Id = dateObj.Id
				band.QueryResult = date
				results = append(results, band)
			}
		}
	}

	// Search for matching locations
	for _, locationObj := range Data.Locations.Index {
		for _, location := range locationObj.Locations {
			if strings.Contains(location, query) {
				band := ResultIDs{}
				band.SearchParam = "Tour Location"
				band.Id = locationObj.Id
				band.QueryResult = location
				results = append(results, band)
			}
		}
	}

	return results
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

func GetResults(ids []ResultIDs, data api.Data) []api.Artist {
	var results []api.Artist
	var usedIds []int

	for i := range ids {
		for _, artist := range data.Artists {
			if ids[i].Id == artist.Id && !IdExists(usedIds, artist.Id) {
				usedIds = append(usedIds, artist.Id)
				results = append(results, artist)
				break
			}
		}
	}

	return results
}

func SuggestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("artist") // Retrieve search query from html form
	results := SearchArtist(query)
	//results := GetResults(ids, Data)
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
