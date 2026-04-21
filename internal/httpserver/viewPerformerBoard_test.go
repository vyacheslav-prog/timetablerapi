package httpserver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"timetablerapi/internal/services"
	"timetablerapi/overview"
)

type overviewStub struct {
}

func (os overviewStub) ViewPerformerBoard(context.Context, string) (*overview.PerformerBoard, error) {
	return &overview.PerformerBoard{}, nil
}

func TestViewPerformerBoardIsError(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/performer-boards/", http.NoBody))
	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Error("expected status code 404, given:", resp.StatusCode)
	}
}

func TestViewPerformerBoardIsSuccess(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{Overview: &overviewStub{}})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/performer-boards/x1", http.NoBody))
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Error("expected status code 200, given:", resp.StatusCode)
	}
}
