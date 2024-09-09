package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	ascii "ascii/ressources"
)

func fileExists(filename string) error {
	_, err := os.Stat(filename)

	return err
}

func main() {
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		path = path[len("/static/"):]
		allowedExtensions := map[string]bool{
			".css":   true,
			".png":   true,
			".jpg":   true,
			".svg":   true,
			".woff2": true,
		}

		ext := filepath.Ext(path)
		if !allowedExtensions[ext] {
			http.ServeFile(w, r, "templates/404.html")
			return
		}

		fullPath := filepath.Join("static", path)
		err := fileExists(fullPath)
		if err == nil {
			http.ServeFile(w, r, fullPath)
		} else {
			http.ServeFile(w, r, "templates/404.html")
			return
		}
	})

	http.HandleFunc("/", ascii.HandleHome)
	http.HandleFunc("/ascii-art", ascii.HandleAsciiArt)
	http.HandleFunc("/authors", ascii.HandleAuthors)
	http.HandleFunc("/about", ascii.HandleAbout)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
