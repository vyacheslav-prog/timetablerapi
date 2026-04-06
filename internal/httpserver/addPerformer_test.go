package httpserver

import (
	"context"
	"testing"
	"timetablerapi/internal/services"
	"timetablerapi/registrar"
)

func TestEmptyAddPerformerRequest(t *testing.T) {
	rg := struct {
		AddPerformer func(context.Context, registrar.Performer) (string, error)
	} {
		AddPerformer: func(_ context.Context, _ registrar.Performer) (string, error) {
			return "", nil
		},
	}
	NewMux(&services.Services{Registrar: rg})
}
