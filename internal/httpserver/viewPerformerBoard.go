package httpserver

import (
	"context"
	"net/http"

	"timetablerapi/overview"
)

type viewPerformerService interface {
	ViewPerformerBoard(context.Context, string) (overview.PerformerBoard, error)
}

func handleViewPerformerBoard(s viewPerformerService, w http.ResponseWriter, r *http.Request) {
	res, ovErr := s.ViewPerformerBoard(r.Context(), r.PathValue("boardId"))
	if ovErr != nil {
		http.Error(w, ovErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, []byte(res.Title()))
}
