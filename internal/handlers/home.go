package handlers

import (
	"groupie-tracker/internal/api"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict access to home page
	if r.Method != "GET" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Load data from API, handle errors if necessary
	data, err := api.LoadData()
	if err != nil {
		http.Error(w, "Could not load API data: ", http.StatusInternalServerError)
		log.Println("Error loading API data: ", err)
	}

	// Create Html template from home.html, handle errors if necessary
	tmpl, err := template.ParseFiles("web/templates/home.html")

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
