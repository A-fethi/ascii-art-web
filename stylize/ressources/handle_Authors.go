package ascii

import (
	"html/template"
	"net/http"
)

func HandleAuthors(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "templates/authors.html")
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
