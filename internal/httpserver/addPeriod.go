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

func handleAddPeriod(_ services.RegistrarService, _ http.ResponseWriter, r *http.Request) {
	var req addPeriodRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print("failed body decode:", err)
	}
}
