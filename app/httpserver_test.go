package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newStubbedServer() *http.ServeMux {
	mux, services := http.NewServeMux(), &services{}
	registerHandlers(mux, services)
	return mux
}

func TestListens8080PortForHttpServer(t *testing.T) {
	s := newStubbedServer()
	req, w := httptest.NewRequest("GET", "http://localhost:8080/", nil), httptest.NewRecorder()
	s.ServeHTTP(w, req)
	resp := w.Result()
	if http.StatusOK != resp.StatusCode {
		t.Errorf("Result for GET / must be [%v], actual is [%v]", http.StatusOK, resp.Status)
	}
}

func TestMissesUnknownPathWith404Status(t *testing.T) {
	s, url := newStubbedServer(), "/some_unknown_path"
	req, w := httptest.NewRequest("GET", url, nil), httptest.NewRecorder()
	s.ServeHTTP(w, req)
	resp := w.Result()
	if expected := http.StatusNotFound; expected != resp.StatusCode {
		t.Errorf("Result for GET [%v] must be [%v], actual is [%v]", url, expected, resp.Status)
	}
}

func TestHandlesGetForPerformerBoard(t *testing.T) {
	s, path := newStubbedServer(), "/performer-boards/1"
	req, w := httptest.NewRequest("GET", path, nil), httptest.NewRecorder()
	s.ServeHTTP(w, req)
	res := w.Result()
	if expected := http.StatusOK; expected != res.StatusCode {
		t.Errorf("Result for GET [%v] must be [%v], actual is [%v]", path, expected, res.Status)
	}
	if contentType, expected := res.Header.Get("Content-Type"), "application/json"; expected != contentType {
		t.Errorf("Content for GET [%v] must be [%v], actual is [%v]", path, expected, contentType)
	}
}
