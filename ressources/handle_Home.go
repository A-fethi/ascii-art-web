package ascii

import (
	"html/template"
	"net/http"
)

// HandleHome serves the home page if the request is valid, or handles errors for invalid requests.
func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		HandleError(w, http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
}
