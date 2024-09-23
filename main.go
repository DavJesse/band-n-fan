package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie-tracker/internal/handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Too Many Arguments")
		return
	}

	http.HandleFunc("/static/", handlers.StaticHandler)
	http.HandleFunc("/", handlers.HomeHandler)               // Serve home page
	http.HandleFunc("/artist/", handlers.ArtistHandler)      // Serve artist.html
	http.HandleFunc("/dates", handlers.DateHandler)          // Serve date.html
	http.HandleFunc("/locations", handlers.LocationsHandler) // Serve location.html
	http.HandleFunc("/relations", handlers.RelationsHandler) // Serve relation.html

	// Start server
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
