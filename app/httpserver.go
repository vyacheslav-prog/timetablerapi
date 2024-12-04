package main

import (
	"fmt"
	"net/http"
)

func newServer(storage *Storage) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, storage.getDashboardTitle())
	})
	mux.HandleFunc("/404", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Page not found")
	})
	return mux
}

func main() {
	storage := newStorage()
	mux := newServer(storage)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error for server startup:", err)
	}
	fmt.Println("Timetablerapi server for http listen 8080 port")
}
