package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func DateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/date.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		log.Println("Failed to load template: ", err)
		return
	}

	// Populate BandName field of Dates struct
	for i := range data.Dates.Index {
		for j := i; j < len(data.Artists); j++ {

			// if data.Dates.Index.Id and data.Artist.Id match, update BandName in dates.Index[i]
			if data.Artists[j].Id == data.Dates.Index[i].Id {
				data.Dates.Index[i].BandName = data.Artists[j].Name
				i++ // Break loop, match found
			}
		}
	}

	err = tmpl.Execute(w, data.Dates)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		log.Println("Failed to execute template", err)
		return
	}
}
