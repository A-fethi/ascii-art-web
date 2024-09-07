package ascii

import (
	"html/template"
	"net/http"
	"strings"
)

func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/index.html",
		"templates/400.html",
		"templates/404.html",
		"templates/500.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "templates/400.html")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")
	slice := strings.Split(text, "\n")
	for i := 0; i < len(slice); i++ {
		for _, char := range slice[i] {
			if char < 32 || char > 126 {
				w.WriteHeader(http.StatusBadRequest)
				http.ServeFile(w, r, "templates/400.html")
				return
			}
		}
	}
	asciiArt := GenerateAsciiArt(slice, banner)
	if asciiArt == "" {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}
	err = tmpl.Execute(w, asciiArt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}
}
