package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	services, err := newServices()
	if err != nil {
		fmt.Println("unable initalization for services:", err)
		os.Exit(1)
	}
	mux := http.NewServeMux()
	registerHandlers(mux, services)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("failed serve for http-server:", err)
	}
	fmt.Println("timetablerapi server for http listen 8080 port")
}

type performerCreatingRequest struct {
	name string
}

func registerHandlers(mux *http.ServeMux, s *services) {
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
	})
	mux.HandleFunc("GET /performer-boards/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, s.overview.ViewPerformerBoard(r.PathValue("boardId")))
	})
	mux.HandleFunc("POST /performers", func(w http.ResponseWriter, r *http.Request) {
		body, readBodyErr := io.ReadAll(r.Body)
		if readBodyErr != nil {
			fmt.Fprint(w, "failed body read:", readBodyErr)
		}
		var data performerCreatingRequest
		if unmarshallErr := json.Unmarshal(body, &data); unmarshallErr != nil {
			fmt.Fprint(w, "failed body decoding:", unmarshallErr)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		s.registrar.AddPerformer(data.name)
		fmt.Fprint(w, "{\"performer_id\": \"a-a-a-a\"}")
	})
}
