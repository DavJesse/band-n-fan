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

	// Populate BandName field of Locations struct
	for i := range data.Locations.Index {
		for j := range data.Artists {
			if data.Artists[j].Id == data.Locations.Index[i].Id {
				data.Locations.Index[i].BandName = data.Artists[j].Name
				break // Break loop when match is found
			}
		}
	}

	// Execute locations template
	err := locationTemplate.Execute(w, data.Locations)
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to execute template:", err)
		return
	}
}
