package httpserver

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"timetablerapi/internal/services"
	"timetablerapi/overview"
)

type overviewStub struct {
}

func (os overviewStub) ViewPerformerBoard(context.Context, string) (*overview.PerformerBoard, error) {
	return overview.NewPerformerBoard("", "", "board"), nil
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
	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("failed body reading:", err)
	}
	bs := string(bb)
	if !strings.Contains(bs, "board") {
		t.Error("performer board must contains into title [board], given body:", bs)
	}
}
