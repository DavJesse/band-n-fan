package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Issue struct {
	StatusCode int
	Problem    string
}

// Initialize variable to hold error message and status codes

var hitch Issue

// Template loader function to allow mocking during tests
var LoadTemplate = func() (*template.Template, error) {
	return template.ParseFiles("web/templates/error.html")
}

func BadRequestHandler(w http.ResponseWriter) {
	tmpl, err := LoadTemplate()
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusBadRequest
	hitch.Problem = "Bad Request!"

	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func InternalServerErrorHandler(w http.ResponseWriter) {
	tmpl, err := LoadTemplate()
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusInternalServerError
	hitch.Problem = "Internal Server Error!"

	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func NotFoundHandler(w http.ResponseWriter) {
	tmpl, err := LoadTemplate()
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusNotFound
	hitch.Problem = "Not Found!"

	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}
