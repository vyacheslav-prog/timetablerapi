package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"timetablerapi/internal/services"
)

type addPeriodRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func handleAddPeriod(s services.RegistrarService, w http.ResponseWriter, r *http.Request) {
	var req addPeriodRequest
	dcdErr := json.NewDecoder(r.Body).Decode(&req)
	if dcdErr != nil {
		log.Print("failed body decode:", dcdErr)
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, []byte(dcdErr.Error()))
		return
	}
	res, regErr := s.AddPeriod(req.From, req.To)
	if regErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, []byte(regErr.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, []byte(res))
}
