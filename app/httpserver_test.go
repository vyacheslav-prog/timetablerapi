package main

import (
	"net/http"
	"testing"
	"time"
)

func TestListens8080PortForHttpServer(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Could not reach to server: %v", err)
	}
	defer resp.Body.Close()
	if http.StatusOK != resp.StatusCode {
		t.Errorf("Result for GET / must be [%v], actual is [%v]", http.StatusOK, resp.Status)
	}
}

func TestMissesUnknownPathWith404Status(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	path := "/some_unknown_path"
	resp, err := http.Get("http://localhost:8080" + path)
	if err != nil {
		t.Fatalf("Could not reach to server: %v", err)
	}
	defer resp.Body.Close()
	if expected := http.StatusNotFound; expected != resp.StatusCode {
		t.Errorf("Result for GET [%v] must be [%v], actual is [%v]", path, expected, resp.Status)
	}
}
