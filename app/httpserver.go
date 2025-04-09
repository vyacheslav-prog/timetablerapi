package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	services, err := newServices()
	if err != nil {
		fmt.Println("Unable initalization for services:", err)
		os.Exit(1)
	}
	mux := http.NewServeMux()
	registerHandlers(mux, services)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error for server startup:", err)
	}
	fmt.Println("Timetablerapi server for http listen 8080 port")
}

func registerHandlers(mux *http.ServeMux, s *services) {
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
	})
	mux.HandleFunc("GET /performer-boards/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, s.overview.ViewPerformerBoard(r.PathValue("boardId")))
	})
}
