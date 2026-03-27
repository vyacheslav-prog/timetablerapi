package httpserver

import (
	"encoding/json"
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
		http.Error(w, dcdErr.Error(), http.StatusBadRequest)
		return
	}
	res, regErr := s.AddPeriod(r.Context(), req.From, req.To)
	if regErr != nil {
		http.Error(w, regErr.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, []byte(res))
}
