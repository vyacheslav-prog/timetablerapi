package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"timetablerapi/internal/services"
)

type addPeriodRequest struct {
	from, to string
}

func handleAddPeriod(s services.RegistrarService, w http.ResponseWriter, r *http.Request) {
	var req addPeriodRequest
	dcdErr := json.NewDecoder(r.Body).Decode(&req)
	if dcdErr != nil {
		log.Print("failed body decode:", dcdErr)
	}
	res, regErr := s.AddPeriod(req.from, req.to)
	if regErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponse(w, []byte(regErr.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, []byte(res))
}
