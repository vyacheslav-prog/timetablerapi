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
	req, rr := httptest.NewRequest("GET", "http://localhost:8080/", nil), httptest.NewRecorder()
	s.ServeHTTP(rr, req)
	resp := rr.Result()
	if expected := http.StatusOK; expected != resp.StatusCode {
		t.Errorf("Result for GET / must be [%v], actual is [%v]", expected, resp.Status)
	}
}

func TestMissesUnknownPathWith404Status(t *testing.T) {
	s, url := newStubbedServer(nil), "/some_unknown_path"
	req, rr := httptest.NewRequest("GET", url, nil), httptest.NewRecorder()
	s.ServeHTTP(rr, req)
	resp := rr.Result()
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

func TestHandlesPostForPerformer(t *testing.T) {
	s := newStubbedServer(nil)
	path := "/performers"
	req, rr := httptest.NewRequest("POST", path, nil), httptest.NewRecorder()
	s.ServeHTTP(rr, req)
	res := rr.Result()
	if expected := http.StatusCreated; expected != res.StatusCode {
		t.Errorf("Result for POST [%v] must have [%v] status, actual is [%v]", path, expected, res.Status)
	}
	if contentType, expected := res.Header.Get("Content-Type"), "application/json"; expected != contentType {
		t.Errorf("Type of content for POST [%v] must be [%v], actual is [%v]", path, expected, contentType)
	}
}
