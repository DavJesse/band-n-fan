package handlers

import (
	"groupie-tracker/internal/api"
	"html/template"
	"log"
	"net/http"
)

var data, err = api.LoadData() // Load data from API
var tmpl *template.Template    // Initialize variable to hold our html templates

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict access to home page, execute error template
	if r.Method != "GET" || r.URL.Path != "/" {
		badRequestHandler(w)

		if r.Method != "GET" {
			log.Println("Bad client request: not GET")
		} else {
			log.Println("Invalid client path")
		}
		return
	}

	// handle errors from API call
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Could not load API data")
		return
	}

	// Create Html template from home.html, handle errors if necessary
	tmpl, err = template.ParseFiles("web/templates/home.html")
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to load home template")
		return
	}

	// Safely execute tmpl, handle errors if necessary
	err = tmpl.Execute(w, data)
	if err != nil {
		internalServerErrorHandler(w)
		log.Println("Failed to execute home template")
		return
	}
}
