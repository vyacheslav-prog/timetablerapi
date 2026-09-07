package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"timetablerapi/registrar"
)

type addTaskService interface {
	AddTask(context.Context, registrar.Task) (string, error)
}

type taskCreatingRequest struct {
	From string `json:"from"`
	Name string `json:"name"`
	To   string `json:"to"`
}

func handleAddTask(s addTaskService, w http.ResponseWriter, r *http.Request) {
	var tcr taskCreatingRequest
	dcdErr := json.NewDecoder(r.Body).Decode(&tcr)
	if dcdErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, regErr := s.AddTask(r.Context(), registrar.Task{From: tcr.From, Name: tcr.Name, To: tcr.To})
	if regErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, []byte(res))
}
