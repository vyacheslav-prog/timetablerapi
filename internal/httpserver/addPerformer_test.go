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

type registrarRepoStub struct {
}

func (s registrarRepoStub) SaveEvent(context.Context, uint) error {
	return nil
}

func (s registrarRepoStub) SaveAndIdentifyLayout(context.Context, string) (string, error) {
	return "", nil
}

func (s registrarRepoStub) SaveAndIdentifyPerformer(context.Context, string) (string, error) {
	return "", nil
}

func (s registrarRepoStub) SaveAndIdentifyTask(context.Context, string, string, string) (string, error) {
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
	registerHandlers(mux, &services.Services{Registrar: registrar.Registrar{Repo: registrarRepoStub{}}})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/performers", strings.NewReader(`{"name":"John"}`)))
	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Error("response must have status 201 on valid request, given:", resp.StatusCode)
	}
}
