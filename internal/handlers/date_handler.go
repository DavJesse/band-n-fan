package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

var (
	dateTemplate          *template.Template
	mockDateTemplate      string
	mockDateTemplateError bool
)

// SetMockDateTemplate sets the content for the mock date template
func SetMockDateTemplate(content string) {
	mockDateTemplate = content
}

// SetMockDateTemplateError sets the error state for mock date template loading
func SetMockDateTemplateError(shouldError bool) {
	mockDateTemplateError = shouldError
}

// loadDateTemplate loads the date template, using a mock if specified
func loadDateTemplate() error {
	if mockDateTemplateError {
		return errors.New("mock date template loading error")
	}
	if mockDateTemplate != "" {
		var err error
		dateTemplate, err = template.New("date").Parse(mockDateTemplate)
		return err
	}
	var err error
	dateTemplate, err = template.ParseFiles("web/templates/date.html")
	return err
}

// DateHandler handles the date endpoint
func DateHandler(w http.ResponseWriter, r *http.Request) {
	if err := loadDateTemplate(); err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to load date template:", err)
		return
	}

	// Populate BandName field of Dates struct
	for i := range data.Dates.Index {
		for j := range data.Artists {
			if data.Artists[j].Id == data.Dates.Index[i].Id {
				data.Dates.Index[i].BandName = data.Artists[j].Name
				break // Break loop when match is found
			}
		}
	}

	// Execute dates template, handle errors if found
	err := dateTemplate.Execute(w, data.Dates)
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to execute template:", err)
		return
	}
}
