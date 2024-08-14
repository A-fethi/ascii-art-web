package main

import (
	"log"
	"net/http"

	ascii "ascii/ressources"
)

func main() {
	http.HandleFunc("/", ascii.HandleHome)
	http.HandleFunc("/ascii-art", ascii.HandleAsciiArt)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
