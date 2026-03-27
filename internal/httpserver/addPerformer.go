package httpserver

import (
	"encoding/json"
	"net/http"

	"timetablerapi/internal/services"
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
	res, regErr := s.AddPerformer(r.Context(), data.Name)
	if regErr != nil {
		http.Error(w, regErr.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, []byte(res))
}
