package main

import (
	"fmt"
	"net/http"
)

func newServer(storage *Storage) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
		}
		fmt.Fprintln(w, storage.getDashboardTitle())
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
