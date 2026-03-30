package httpserver

import (
	"encoding/json"
	"net/http"

	"timetablerapi/internal/services"
	"timetablerapi/registrar"
)

type performerCreatingRequest struct {
	Name string `json:"name"`
}

func handleAddPerformer(s services.RegistrarService, w http.ResponseWriter, r *http.Request) {
	var data performerCreatingRequest
	dcdErr := json.NewDecoder(r.Body).Decode(&data)
	if dcdErr != nil {
		http.Error(w, dcdErr.Error(), http.StatusBadRequest)
		return
	}
	prf := registrar.Performer{Name: data.Name}
	res, regErr := s.AddPerformer(r.Context(), prf)
	if regErr != nil {
		http.Error(w, regErr.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, []byte(res))
}
