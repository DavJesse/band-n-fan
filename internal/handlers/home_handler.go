package handlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/internal/api"
)

var (
	Data, err         = api.LoadData() // Load data from API
	tmpl              *template.Template
	mockTemplateError bool
	mockTemplate      string
)

func SetMockHomeTemplate(content string) {
	mockTemplate = content
}

func SetMockHomeTemplateError(shouldError bool) {
	mockTemplateError = shouldError
}

func loadHomeTemplate() error {
	if mockTemplateError {
		return errors.New("mock template loading error")
	}
	if mockTemplate != "" {
		var err error
		tmpl, err = template.New("home").Parse(mockTemplate)
		return err
	}
	var err error
	tmpl, err = template.ParseFiles("web/templates/home.html")
	return err
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" || r.URL.Path != "/" {
		BadRequestHandler(w)
		log.Println("Home handler bad error")
		return
	}

	if err := loadHomeTemplate(); err != nil {
		log.Println("Failed to load home template:", err)
		InternalServerErrorHandler(w)
		return
	}

	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Could not load API data:", err)
		return
	}

	// Safely execute tmpl, handle errors if necessary
	err = tmpl.Execute(w, Data)
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Failed to execute home template")
		return
	}
}
