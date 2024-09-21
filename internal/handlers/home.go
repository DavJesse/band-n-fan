package handlers

import (
	"groupie-tracker/internal/api"
	"html/template"
	"log"
	"net/http"
)

var data, err = api.LoadData() // Load data from API
var tmpl *template.Template

type Issue struct {
	StatusCode int
	Problem    string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict access to home page
	if r.Method != "GET" {
		tmpl, err = template.ParseFiles("web/templates/error.html")
		Issue.StatusCode = http.StatusBadRequest

		tmpl.Execute()
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// handle errors from API call
	if err != nil {
		http.Error(w, "Could not load API data: ", http.StatusInternalServerError)
		log.Println("Error loading API data: ", err)
	}

	// Create Html template from home.html, handle errors if necessary
	tmpl, err = template.ParseFiles("web/templates/home.html")

	if err != nil {
		http.Error(w, "Could not load home template", http.StatusInternalServerError)
		log.Println("Error parsing template: ", err)
		return
	}

	// Safely execute tmpl, handle errors if necessary
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Could not execute home template", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
		return
	}
}
