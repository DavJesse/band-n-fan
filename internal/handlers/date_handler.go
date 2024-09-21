package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func DateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/date.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		log.Println("Failed to load template: ", err)
		return
	}

	err = tmpl.Execute(w, data.Dates)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		log.Println("Failed to execute template", err)
		return
	}
}
