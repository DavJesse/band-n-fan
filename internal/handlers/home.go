package handlers

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	//"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//path := filepath.Join("web", "templates", "home.go")
	if r.Method != "GET" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	//fmt.Println(path)

	tmpl, err := template.ParseFiles("C:\\Users\\pc\\OneDrive\\Desktop\\module\\groupie-tracker\\web\\templates\\home.html")

	if err != nil {
		http.Error(w, "Could not load home template", http.StatusInternalServerError)
		log.Println("Error parsing template: ", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Could not execute home template", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
		return
	}
}
