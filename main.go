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

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve home page
	http.HandleFunc("/", handlers.HomeHandler)

	// Start server
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
