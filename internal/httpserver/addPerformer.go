package httpserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"timetablerapi/internal/services"
)

type performerCreatingRequest struct {
	Name string `json:"name"`
}

func handleAddPerformer(s services.RegistrarService, w http.ResponseWriter, r *http.Request) {
	body, readBodyErr := io.ReadAll(r.Body)
	if readBodyErr != nil {
		log.Print("failed body read:", readBodyErr)
	}
	var data performerCreatingRequest
	if unmarshallErr := json.Unmarshal(body, &data); unmarshallErr != nil {
		log.Print("failed body decoding:", unmarshallErr)
	}
	res, regErr := s.AddPerformer(data.Name)
	if regErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, []byte(regErr.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, []byte(res))
}
