package ascii

import "net/http"

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/about.html")
}