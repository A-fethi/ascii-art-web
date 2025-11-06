package ascii

import (
	"html/template"
	"net/http"
)

// HandleAbout serves the "authors" page template or returns a 500 error if something goes wrong.
func HandleAuthors(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/authors.html")
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
