package ascii

import (
	"net/http"
	"text/template"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Error 404: Template not found", http.StatusNotFound)
		return
	}

	q := r.URL.Query().Get("error")
	println(q)
	if r.URL.Query().Get("error") == "true" {
		http.Error(w, "Error 500: Internal server error!", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Error 400: Bad Request!", http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Error 404: Page Not Found!", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, nil)
}
