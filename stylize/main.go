package main

import (
	"log"
	"net/http"

	ascii "ascii/ressources"
)

func main() {
	http.HandleFunc("/static/", ascii.HandleStatic)
	http.HandleFunc("/", ascii.HandleHome)
	http.HandleFunc("/ascii-art", ascii.HandleAsciiArt)
	http.HandleFunc("/authors", ascii.HandleAuthors)
	http.HandleFunc("/about", ascii.HandleAbout)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
