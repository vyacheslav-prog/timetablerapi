package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	srvAddr              = ":8080"
	srvReadHeaderTimeout = 15 * time.Second
	srvReadTimeout       = 10 * time.Second
	srvWriteTimeout      = 10 * time.Second
	srvIdleTimeout       = 30 * time.Second
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
	httpSrv := &http.Server{
		Addr:              srvAddr,
		ReadHeaderTimeout: srvReadHeaderTimeout,
		ReadTimeout:       srvReadTimeout,
		WriteTimeout:      srvWriteTimeout,
		IdleTimeout:       srvIdleTimeout,
	}
	if serveErr := httpSrv.ListenAndServe(); serveErr != nil {
		log.Println("failed serve for http-server:", serveErr)
	}
	log.Println("timetablerapi server for http listen 8080 port")
}

type performerCreatingRequest struct {
	Name string `json:"name"`
}

func registerHandlers(mux *http.ServeMux, s *services) {
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
	})
	mux.HandleFunc("GET /performer-boards/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		res, ovErr := s.overview.ViewPerformerBoard(r.Context(), r.PathValue("boardId"))
		if ovErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, respErr := w.Write([]byte(ovErr.Error()))
			if respErr != nil {
				log.Print("can not to write a response:", respErr)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, respErr := fmt.Fprint(w, res)
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
		res, regErr := s.registrar.AddPerformer(data.Name)
		if regErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, respErr := w.Write([]byte(regErr.Error()))
			if respErr != nil {
				log.Print("can not to write a response:", respErr)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, respErr := fmt.Fprint(w, res)
		if respErr != nil {
			log.Print("can not to write a response:", respErr)
		}
	})
}
