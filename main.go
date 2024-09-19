package main

import (
	"fmt"
	"groupie-tracker/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Too Many Arguments")
		return
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static")))) // Serve static files
	http.HandleFunc("/", handlers.HomeHandler)                                                     // Serve home page
	http.HandleFunc("/artist/", handlers.ArtistHandler)                                         // Serve artist.html

	// Start server
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
