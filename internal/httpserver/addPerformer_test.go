package httpserver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"timetablerapi/internal/services"
	"timetablerapi/registrar"
)

type registrarStub struct {
	result string
}

func (rs registrarStub) AddPerformer(context.Context, registrar.Performer) (string, error) {
	return rs.result, nil
}

func (rs registrarStub) AddPeriod(context.Context, string, string) (string, error) {
	return "", nil
}

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
	registerHandlers(mux, &services.Services{Registrar: registrarStub{"ok"}})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/performers", strings.NewReader(`{"name":"John"}`)))
	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Error("response must have status 201 on valid request, given:", resp.StatusCode)
	}
}
