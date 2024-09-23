package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

var (
	relationTemplate         *template.Template
	mockRelationTemplate     string
	mockRelationTemplateError bool
)

func SetMockRelationTemplate(content string) {
	mockRelationTemplate = content
}

func SetMockRelationTemplateError(shouldError bool) {
	mockRelationTemplateError = shouldError
}

func loadRelationTemplate() error {
	if mockRelationTemplateError {
		return errors.New("mock relation template loading error")
	}
	if mockRelationTemplate != "" {
		var err error
		relationTemplate, err = template.New("relation").Parse(mockRelationTemplate)
		return err
	}
	var err error
	relationTemplate, err = template.ParseFiles("web/templates/relation.html")
	return err
}

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	if err := loadRelationTemplate(); err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to load relation template:", err)
		return
	}

	// Populate BandName field of Relations struct
	for i := range data.Relations.Index {
		for j := range data.Artists {
			if data.Artists[j].Id == data.Relations.Index[i].Id {
				data.Relations.Index[i].BandName = data.Artists[j].Name
				break // Break loop when match is found
			}
		}
	}

	// Execute relations template
	err := relationTemplate.Execute(w, data.Relations)
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to execute template:", err)
		return
	}
}
