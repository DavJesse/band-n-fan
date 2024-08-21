package handlers

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	//"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict access to home page
	if r.Method != "GET" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Create tamplate from home.html, handle errors if necessary
	tmpl, err := template.ParseFiles("web/templates/home.html")

	if err != nil {
		http.Error(w, "Could not load home template", http.StatusInternalServerError)
		log.Println("Error parsing template: ", err)
		return
	}

	// Safely execute template, handle errors if necessary
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Could not execute home template", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
		return
	}
}
