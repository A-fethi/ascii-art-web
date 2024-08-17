package ascii

import (
	"net/http"
	"strings"
	"text/template"
)

func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Error 404: Template not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Error 400: Bad request!", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if len(text) > 1000 {
		http.Error(w, "Error 403: Input length must be less than 1000", http.StatusForbidden)
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Error 404: Banner not found!", http.StatusNotFound)
		return
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")
	slice := strings.Split(text, "\n")
	asciiArt := GenerateAsciiArt(slice, banner)

	tmpl.Execute(w, asciiArt)
}
