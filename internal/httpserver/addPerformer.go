package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"timetablerapi/registrar"
)

type addPerformerService interface {
	AddPerformer(context.Context, registrar.Performer) (string, error)
}

type performerCreatingRequest struct {
	Name string `json:"name"`
}

func handleAddPerformer(s addPerformerService, w http.ResponseWriter, r *http.Request) {
	var pcr performerCreatingRequest
	dcdErr := json.NewDecoder(r.Body).Decode(&pcr)
	if dcdErr != nil {
		http.Error(w, dcdErr.Error(), http.StatusBadRequest)
		return
	}
	prf := registrar.Performer{Name: pcr.Name}
	res, regErr := s.AddPerformer(r.Context(), prf)
	if regErr != nil {
		http.Error(w, regErr.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, []byte(res))
}
