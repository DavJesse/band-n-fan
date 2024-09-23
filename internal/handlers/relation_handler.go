package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err = template.ParseFiles("web/templates/relation.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to load template: ", err)
		return
	}

	// Populate BandName field of Relations struct
	for i := range data.Relations.Index {
		for j := i; j < len(data.Artists); j++ {
			// if data.Relations.Index.Id and data.Artist.Id match, update BandName in dates.Index[i]
			if data.Artists[j].Id == data.Relations.Index[i].Id {
				data.Relations.Index[i].BandName = data.Artists[j].Name
				i++ // Break loop, match found
			}
		}
	}

	// Execute locations template
	err = tmpl.Execute(w, data.Relations)
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to execute template", err)
		return
	}
}
