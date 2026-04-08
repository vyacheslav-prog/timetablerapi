package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"timetablerapi/internal/services"
)

func TestEmptyAddPerformerRequest(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest("POST", "/performers", nil))
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Error("response must have status 400 on empty body, given:", resp.StatusCode)
	}
}
