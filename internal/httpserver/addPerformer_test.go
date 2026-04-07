package httpserver

import (
	"context"
	"net/http"
	"testing"
	"timetablerapi/internal/services"
	"timetablerapi/registrar"
)

func TestEmptyAddPerformerRequest(t *testing.T) {
	rg := struct {
		AddPerformer func(context.Context, registrar.Performer) (string, error)
	}{
		AddPerformer: func(_ context.Context, _ registrar.Performer) (string, error) {
			return "", nil
		},
	}
	mux := http.NewServeMux()
	registerHandlers(mux, &services.Services{Registrar: rg})
	req := http.NewRequest("POST", "/performers")
	mux.ServeHTTP()
}
