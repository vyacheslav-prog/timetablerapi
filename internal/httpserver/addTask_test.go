package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"timetablerapi/internal/services"
)

func TestAddTaskIsError(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/tasks", http.NoBody))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Error("response must have 400 status for empty body, given:", resp.Status)
	}
}

func TestAddTaskIsSuccess(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{Registrar: services.RegistrarStub{}})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/tasks", nil))
	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Error("response must have 201 status, given:", resp.Status)
	}
}
