package ascii

import (
	"net/http"
	"html/template"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
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

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
    	http.ServeFile(w, r, "templates/400.html")
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
    	http.ServeFile(w, r, "templates/404.html")
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
    	http.ServeFile(w, r, "templates/500.html")
		return
	}
}
