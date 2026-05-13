package httpserver

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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
	registerHandlers(mux, &services.Services{Registrar: services.RegistrarStub{Result: "ok"}})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"name":"do it","from":"08:00","to":"08:30"}`)))
	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Error("response must have 201 status, given:", resp.Status)
	}
	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("can not read a response:", err)
	}
	if !strings.Contains(string(bb), "ok") {
		t.Error("response must contain result from registrar")
	}
}
