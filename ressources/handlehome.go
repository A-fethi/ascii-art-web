package ascii

import (
	"fmt"
	"net/http"
	"text/template"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
	if r.URL.Path != "/" {
		fmt.Fprint(w, "Page Not Found ", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, nil)
}
