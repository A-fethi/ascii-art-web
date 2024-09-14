package main

import (
	"log"
	"net/http"

	ascii "ascii/ressources"
)

// main sets up HTTP route handlers and starts the web server on port 8080.
func main() {
	http.HandleFunc("/static/", ascii.HandleStatic)
	http.HandleFunc("/", ascii.HandleHome)
	http.HandleFunc("/ascii-art", ascii.HandleAsciiArt)
	http.HandleFunc("/authors", ascii.HandleAuthors)
	http.HandleFunc("/about", ascii.HandleAbout)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
