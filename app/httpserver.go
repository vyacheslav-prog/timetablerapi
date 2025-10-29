package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	services, initErr := newServices(ctx)
	if initErr != nil {
		log.Println("unable initialization for services:", initErr)
		os.Exit(1)
	}
	mux := http.NewServeMux()
	registerHandlers(mux, services)
	if serveErr := http.ListenAndServe(":8080", mux); serveErr != nil {
		log.Println("failed serve for http-server:", serveErr)
	}
	log.Println("timetablerapi server for http listen 8080 port")
}

type performerCreatingRequest struct {
	Name string
}

func registerHandlers(mux *http.ServeMux, s *services) {
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
	})
	mux.HandleFunc("GET /performer-boards/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, respErr := fmt.Fprint(w, s.overview.ViewPerformerBoard(r.PathValue("boardId")))
		if respErr != nil {
			log.Print("can not to write a response:", respErr)
		}
	})
	mux.HandleFunc("POST /performers", func(w http.ResponseWriter, r *http.Request) {
		body, readBodyErr := io.ReadAll(r.Body)
		if readBodyErr != nil {
			log.Print("failed body read:", readBodyErr)
		}
		var data performerCreatingRequest
		if unmarshallErr := json.Unmarshal(body, &data); unmarshallErr != nil {
			log.Print("failed body decoding:", unmarshallErr)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, respErr := fmt.Fprint(w, s.registrar.AddPerformer(data.Name))
		if respErr != nil {
			log.Print("can not to write a response:", respErr)
		}
	})
}
