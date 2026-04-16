package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"timetablerapi/internal/services"
)

func TestViewPerformerBoardIsError(t *testing.T) {
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{})
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/performer-boards/", http.NoBody))
	resp := w.Result()
	t.Log("status code of view performer board:", resp.StatusCode)
}
