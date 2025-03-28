package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newStubbedServer(s *services) *http.ServeMux {
	mux := http.NewServeMux()
	if s != nil {
		registerHandlers(mux, s)
	} else {
		registerHandlers(mux, &services{})
	}
	return mux
}

func TestListens8080PortForHttpServer(t *testing.T) {
	s := newStubbedServer(nil)
	req, w := httptest.NewRequest("GET", "http://localhost:8080/", nil), httptest.NewRecorder()
	s.ServeHTTP(w, req)
	resp := w.Result()
	if http.StatusOK != resp.StatusCode {
		t.Errorf("Result for GET / must be [%v], actual is [%v]", http.StatusOK, resp.Status)
	}
}

func TestMissesUnknownPathWith404Status(t *testing.T) {
	s, url := newStubbedServer(nil), "/some_unknown_path"
	req, w := httptest.NewRequest("GET", url, nil), httptest.NewRecorder()
	s.ServeHTTP(w, req)
	resp := w.Result()
	if expected := http.StatusNotFound; expected != resp.StatusCode {
		t.Errorf("Result for GET [%v] must be [%v], actual is [%v]", url, expected, resp.Status)
	}
}

type StubOverviewService struct{}

func (sos StubOverviewService) ViewPerformerBoard(id string) string {
	return id
}

func TestHandlesGetForPerformerBoard(t *testing.T) {
	s := newStubbedServer(&services{overview: StubOverviewService{}})
	path := "/performer-boards/1"
	req, rr := httptest.NewRequest("GET", path, nil), httptest.NewRecorder()
	s.ServeHTTP(rr, req)
	res := rr.Result()
	if expected := http.StatusOK; expected != res.StatusCode {
		t.Errorf("Result for GET [%v] must be [%v], actual is [%v]", path, expected, res.Status)
	}
	if contentType, expected := res.Header.Get("Content-Type"), "application/json"; expected != contentType {
		t.Errorf("Type of content for GET [%v] must be [%v], actual is [%v]", path, expected, contentType)
	}
	if content := rr.Body.String(); content != "1" {
		t.Errorf("Content for GET [%v] must be 1, actual is [%v]", path, content)
	}
}
