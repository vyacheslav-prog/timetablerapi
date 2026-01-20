package httpserver

import (
	"log"
	"net/http"

	"timetablerapi/internal/services"
)

func NewMux(srvs *services.Services) {
	mux := http.NewServeMux()
	registerHandlers(mux, srvs)
}

func registerHandlers(mux *http.ServeMux, s *services.Services) {
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		handleHome(w, r)
	})
	mux.HandleFunc("GET /performer-boards/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		handleViewPerformerBoard(s.Overview, w, r)
	})
	mux.HandleFunc("POST /performers", func(w http.ResponseWriter, r *http.Request) {
		handleAddPerformer(s.Registrar, w, r)
	})
	mux.HandleFunc("POST /periods", func(w http.ResponseWriter, r *http.Request) {
		handleAddPeriod(s.Registrar, w, r)
	})
}

func writeResponse(w http.ResponseWriter, b []byte) {
	_, respErr := w.Write(b)
	if respErr != nil {
		log.Print("can not to write a response:", respErr)
	}
}
