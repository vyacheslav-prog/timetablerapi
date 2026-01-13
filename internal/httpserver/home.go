package httpserver

import "net/http"

func handleHome(w http.ResponseWriter, _ *http.Request) {
	writeResponse(w, []byte("ok"))
}
