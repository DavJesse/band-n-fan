package handlers

import (
	"groupie-tracker/internal/api"
	"html/template"
	"net/http"
)

type Issue struct {
	StatusCode int
	Problem    string
}

var data, err = api.LoadData() // Load data from API
var tmpl *template.Template
var hitch Issue

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict access to home page, execute error template
	if r.Method != "GET" {
		badRequestHandler(w)
		return
	}

	// handle errors from API call
	if err != nil {
		internalServerErrorHandler(w)
		return
	}

	// Create Html template from home.html, handle errors if necessary
	tmpl, err = template.ParseFiles("web/templates/home.html")
	if err != nil {
		internalServerErrorHandler(w)
		return
	}

	// Safely execute tmpl, handle errors if necessary
	err = tmpl.Execute(w, data)
	if err != nil {
		internalServerErrorHandler(w)
		return
	}
}
