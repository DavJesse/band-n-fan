package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/location.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to load template: ", err)
		return
	}

	// Populate BandName field of Dates struct
	for i := range data.Locations.Index {
		for j := i; j < len(data.Artists); j++ {

			// if data.Dates.Index.Id and data.Artist.Id match, update BandName in dates.Index[i]
			if data.Artists[j].Id == data.Locations.Index[i].Id {
				data.Locations.Index[i].BandName = data.Artists[j].Name
				i++ // Break loop, match found
			}
		}
	}

	// Execute locations template
	err = tmpl.Execute(w, data.Locations)
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to execute template", err)
		return
	}
}
