package main

import (
	"log"
	"net/http"

	ascii "ascii/ressources"
)

func main() {
	// http.HandleFunc("/static/css/", func(w http.ResponseWriter, r *http.Request) {
	// 	file := r.URL.Path[len("/css/"):]

	// 	if file != "style.css" {
	// 		fmt.Fprintf(w, "sir bhalk")
	// 		return
	// 	}
	// 	http.ServeFile(w, r, "./static/css/"+file)
	// })
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", ascii.HandleHome)
	// mux.HandleFunc("/ascii-art", ascii.HandleAsciiArt)
	http.HandleFunc("/", ascii.HandleHome)
	http.HandleFunc("/ascii-art", ascii.HandleAsciiArt)
	http.HandleFunc("/authors", ascii.HandleAuthors)
	// http.HandleFunc("/static", ascii.HandleStatic)
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
