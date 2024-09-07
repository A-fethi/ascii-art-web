package ascii

import "net/http"

func HandleAuthors(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/authors.html")
}
