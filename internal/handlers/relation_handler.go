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
	for i := range Data.Relations.Index {
		for j := i; j < len(Data.Artists); j++ {
			// if data.Relations.Index.Id and data.Artist.Id match, update BandName in dates.Index[i]
			if Data.Artists[j].Id == Data.Relations.Index[i].Id {
				Data.Relations.Index[i].BandName = Data.Artists[j].Name
				i++ // Break loop, match found
			}
		}
	}

	// Execute locations template
	err = tmpl.Execute(w, Data.Relations)
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to execute template", err)
		return
	}
}
