package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "static/index.html"
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, "./"+path)
}

func Content(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "static/form.html"
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, "./"+path)
}

func main() {
	log.Println("Starting the server")

	http.HandleFunc("/", Home)
	http.HandleFunc("/ascii-art", Content)

	log.Println("Started on port", "http://localhost:8080")
	fmt.Println("To close connection CTRL-C")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
