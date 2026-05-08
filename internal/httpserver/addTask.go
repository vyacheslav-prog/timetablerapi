package httpserver

import (
	"encoding/json"
	"net/http"
	"timetablerapi/internal/services"
	"timetablerapi/registrar"
)

type taskCreatingRequest struct {
	Name string `json:"name"`
}

func handleAddTask(s services.RegistrarService, w http.ResponseWriter, r *http.Request) {
	var tcr taskCreatingRequest
	dcdErr := json.NewDecoder(r.Body).Decode(&tcr)
	if dcdErr != nil {
		w.WriteHeader(400)
		return
	}
	res, regErr := s.AddTask(r.Context(), registrar.Task{Name: tcr.Name})
	if regErr != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
	writeResponse(w, []byte(res))
}
