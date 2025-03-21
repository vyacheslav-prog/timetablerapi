package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux, services := http.NewServeMux(), newServices()
	registerHandlers(mux, services)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error for server startup:", err)
	}
	fmt.Println("Timetablerapi server for http listen 8080 port")
}

func registerHandlers(mux *http.ServeMux, s *services) {
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, s.storage.getDashboardTitle())
	})
	mux.HandleFunc("GET /performer-boards/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, s.overview.viewPerformerBoard(r.PathValue("boardId")))
	})
}
