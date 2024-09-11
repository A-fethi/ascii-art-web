package ascii

import (
	"html/template"
	"net/http"
	"strings"
)

func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	// files := []string{
	// 	"templates/index.html",
	// 	"templates/400.html",
	// 	"templates/404.html",
	// 	"templates/500.html",
	// }
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		HandleError(w, http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	if text == "" {
		HandleError(w, http.StatusBadRequest)
		return
	}
	text = strings.ReplaceAll(text, "\r\n", "\n")
	if text[0] == '\n' {
		text = "\n" + text
	}

	slice := strings.Split(text, "\n")
	if strings.ReplaceAll(text, "\n", "") == "" && len(text) > 1 {
		slice = slice[1:]
	}
	for i := 0; i < len(slice); i++ {
		for _, char := range slice[i] {
			if char < 32 || char > 126 {
				HandleError(w, http.StatusBadRequest)
				return
			}
		}
	}
	asciiArt := GenerateAsciiArt(slice, banner)
	if asciiArt == "" {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, asciiArt)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
}
