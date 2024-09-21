package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func badRequestHandler(w http.ResponseWriter) {
	tmpl, err = template.ParseFiles("web/templates/error.html")

	// Use writer to render error if error page unavailable
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusBadRequest
	hitch.Problem = "Bad Request!"

	// Render template in html, handle errors if necesary
	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func internalServerErrorHandler(w http.ResponseWriter) {
	tmpl, err = template.ParseFiles("web/templates/error.html")

	// Use writer to render error if error page unavailable
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusInternalServerError
	hitch.Problem = "Internal Server Error!"

	// Render template in html
	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func notFoundHandler(w http.ResponseWriter) {
	tmpl, err = template.ParseFiles("web/templates/error.html")

	// Use writer to render error if error page unavailable
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	hitch.StatusCode = http.StatusNotFound
	hitch.Problem = "Not Found!"

	// Render template in html
	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}