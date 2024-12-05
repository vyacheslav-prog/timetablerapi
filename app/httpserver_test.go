package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListens8080PortForHttpServer(t *testing.T) {
	s := newServer()
	req, w := httptest.NewRequest("GET", "http://localhost:8080/"), httptest.NewRecorder()
	s.ServeHTTP(w, req)
	resp := w.Result()
	if http.StatusOK != resp.StatusCode {
		t.Errorf("Result for GET / must be [%v], actual is [%v]", http.StatusOK, resp.Status)
	}
}

func TestMissesUnknownPathWith404Status(t *testing.T) {
	s, url := newServer(), "http://localhost:8080/some_unknown_path"
	req, w := httptest.NewRequest("GET", url), httptest.NewRecorder()
	s.ServeHTTP(w, req)
	resp := w.Result()
	if expected := http.StatusNotFound; expected != resp.StatusCode {
		t.Errorf("Result for GET [%v] must be [%v], actual is [%v]", url, expected, resp.Status)
	}
}
