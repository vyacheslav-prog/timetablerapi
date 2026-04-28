package httpserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"timetablerapi/internal/services"
)

func TestAddPerformerIsError(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/performers", http.NoBody))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Error("response must have status 400 on empty body, given:", resp.StatusCode)
	}
}

func TestAddPerformerIsSuccess(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{Registrar: services.RegistrarStub{Result: "ok"}})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/performers", strings.NewReader(`{"name":"John"}`)))
	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Error("response must have status 201 on valid request, given:", resp.StatusCode)
	}
}
