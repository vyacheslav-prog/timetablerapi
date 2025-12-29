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

func handleViewPerformerBoard(s services.OverviewService, w http.ResponseWriter, r *http.Request) {
	res, ovErr := s.ViewPerformerBoard(r.Context(), r.PathValue("boardId"))
	if ovErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, []byte(ovErr.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, []byte(res))
}

func registerHandlers(mux *http.ServeMux, s *services.Services) {
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		writeResponse(w, []byte("ok"))
	})
	mux.HandleFunc("GET /performer-boards/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		handleViewPerformerBoard(s.Overview, w, r)
	})
	mux.HandleFunc("POST /performers", func(w http.ResponseWriter, r *http.Request) {
		handleAddPerformer(s.Registrar, w, r)
	})
}

func writeResponse(w http.ResponseWriter, b []byte) {
	_, respErr := w.Write(b)
	if respErr != nil {
		log.Print("can not to write a response:", respErr)
	}
}
