package httpserver

import (
	"encoding/json"
	"net/http"
	"timetablerapi/internal/services"
)

type taskCreatingRequest struct {
	Name string `json:"name"`
}

func handleAddTask(_ services.RegistrarService, w http.ResponseWriter, r *http.Request) {
	var tcr taskCreatingRequest
	dcdErr := json.NewDecoder(r.Body).Decode(&tcr)
	if dcdErr != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(201)
}
