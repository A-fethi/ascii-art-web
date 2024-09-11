package ascii

import (
	"html/template"
	"net/http"
)

func HandleError(w http.ResponseWriter, StatusCodes int) {
	templates := map[int]string{
		400: "templates/400.html",
		404: "templates/404.html",
		500: "templates/500.html",
	}
	nametempl, exicte := templates[StatusCodes]
	if !exicte {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(nametempl)
	if err != nil {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(StatusCodes)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
		return
	}
}
