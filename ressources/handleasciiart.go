package ascii

import (
	"net/http"
	"strings"
	"text/template"
)

func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if len(text) > 1000 {
		http.Error(w, "Input length must be less than 1000", http.StatusForbidden)
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")
	slice := strings.Split(text, "\n")
	asciiArt := GenerateAsciiArt(slice, banner)

	// w.Header().Set("Content-Type", "text/plain")
	// w.Write([]byte(asciiArt))
	tmpl.Execute(w, asciiArt)
}
