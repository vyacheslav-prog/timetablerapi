package main

import (
	"fmt"
	"net/http"
)

func main() {
	storage := newStorage()
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, storage.getDashboardTitle())
	})
	fmt.Println("Timetablerapi server for http listen 8080 port")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error for server startup:", err)
	}
}
