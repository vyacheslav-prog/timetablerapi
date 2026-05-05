package httpserver

import "net/http"

func handleAddTask (w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(400)
}
