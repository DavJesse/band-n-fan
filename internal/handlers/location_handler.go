package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

var (
	locationTemplate         *template.Template
	mockLocationTemplate     string
	mockLocationTemplateError bool
)

func SetMockLocationTemplate(content string) {
	mockLocationTemplate = content
}

func SetMockLocationTemplateError(shouldError bool) {
	mockLocationTemplateError = shouldError
}

func loadLocationTemplate() error {
	if mockLocationTemplateError {
		return errors.New("mock location template loading error")
	}
	if mockLocationTemplate != "" {
		var err error
		locationTemplate, err = template.New("location").Parse(mockLocationTemplate)
		return err
	}
	var err error
	locationTemplate, err = template.ParseFiles("web/templates/location.html")
	return err
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	if err := loadLocationTemplate(); err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to load location template:", err)
		return
	}

	// Populate BandName field of Dates struct
	for i := range Data.Locations.Index {
		for j := i; j < len(Data.Artists); j++ {
			// if data.Dates.Index.Id and data.Artist.Id match, update BandName in dates.Index[i]
			if Data.Artists[j].Id == Data.Locations.Index[i].Id {
				Data.Locations.Index[i].BandName = Data.Artists[j].Name
				i++ // Break loop, match found
			}
		}
	}

	// Execute locations template
	err = tmpl.Execute(w, Data.Locations)
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to execute template:", err)
		return
	}
}
