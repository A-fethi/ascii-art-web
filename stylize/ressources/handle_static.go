package ascii

import "net/http"

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/404.html")
}