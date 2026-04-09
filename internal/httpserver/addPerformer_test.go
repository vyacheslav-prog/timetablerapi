package httpserver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"timetablerapi/internal/services"
	"timetablerapi/registrar"
)

type RegistrarStub struct {}

func (rs RegistrarStub) AddPerformer(context.Context, registrar.Performer) (string, error) {
	return "", nil
}

func TestAddPerformerIsError(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest("POST", "/performers", nil))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Error("response must have status 400 on empty body, given:", resp.StatusCode)
	}
}

func TestAddPerformerIsSuccess(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{Registrar: RegistrarStub{}})
}
