package handlers

import (
	"errors"
	"groupie-tracker/internal/api"
	"html/template"
	"log"
	"net/http"
)

var (
	data, err         = api.LoadData() // Load data from API
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
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := loadHomeTemplate(); err != nil {
		log.Println("Failed to load home template:", err)
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	if err != nil {
		log.Println("Could not load API data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Failed to execute home template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
