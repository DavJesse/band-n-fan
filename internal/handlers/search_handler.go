package handlers

import (
	"groupie-tracker/internal/api"
	"net/http"
	"strconv"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Response) {

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
