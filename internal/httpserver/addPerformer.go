package httpserver

import (
	"encoding/json"
	"log"
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
		log.Print("failed body decode:", dcdErr)
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, []byte(dcdErr.Error()))
		return
	}
	res, regErr := s.AddPerformer(r.Context(), data.Name)
	if regErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, []byte(regErr.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, []byte(res))
}
