package main

import (
	"fmt"
	"net/http"
)

func newServer(services *services) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, services.storage.getDashboardTitle())
	})
	mux.HandleFunc("GET /performer-boards/{boardId}", func(w http.ResponseWriter, r *http.Request) {
	})
	return mux
}

func main() {
	services := newServices()
	mux := newServer(services)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error for server startup:", err)
	}
	fmt.Println("Timetablerapi server for http listen 8080 port")
}
