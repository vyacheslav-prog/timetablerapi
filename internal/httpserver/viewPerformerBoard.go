package httpserver

import (
	"net/http"

	"timetablerapi/internal/services"
)

func handleViewPerformerBoard(s services.OverviewService, w http.ResponseWriter, r *http.Request) {
	res, ovErr := s.ViewPerformerBoard(r.Context(), r.PathValue("boardId"))
	if ovErr != nil {
		http.Error(w, ovErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, []byte(res))
}
