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
		log.Println("Failed to execute template:", err)
		return
	}
}
