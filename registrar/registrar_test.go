package registrar

import (
	"context"
	"testing"
)

type stubRepository struct { }

func (r stubRepository) SaveAndIdentifyLayout(context.Context, string) (string, error) {
	return "", nil
}

func (r stubRepository) SaveAndIdentifyPerformer(context.Context, string) (string, error) {
	return "", nil
}

func (r stubRepository) SaveAndIdentifyTask(context.Context, string, string, string) (string, error) {
	return "", nil
}

func TestNoFiresEventWhenPerformerAddingIsFail(t *testing.T) {
	reg := &Registrar{Repo: stubRepository{}}
	_, err := reg.AddPerformer(t.Context(), Performer{})
	if err == nil {
		t.Error("failed try must return error, given nil")
		return
	}
}
