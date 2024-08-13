package main

import (
	"html/template"
	"log"
	"net/http"

	ascii "ascii/f"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ascii-art", handleAsciiArt)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	asciiArt, err := ascii.GenerateAsciiArt(text, banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(asciiArt))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func Home(w http.ResponseWriter, r *http.Request) {
// 	path := r.URL.Path
// 	if path == "/" {
// 		path = "static/index.html"
// 	}
// 	w.Header().Set("Content-Type", "text/html")
// 	http.ServeFile(w, r, "./"+path)
// }

// func main() {
// 	log.Println("Starting the server")

// 	http.HandleFunc("/", Home)

// 	log.Println("Started on port", "http://localhost:8080")
// 	fmt.Println("To close connection CTRL-C")

// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
