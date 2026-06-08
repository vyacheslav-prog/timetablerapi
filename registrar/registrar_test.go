package registrar

import (
	"context"
	"errors"
	"testing"
)

type stubRepository struct {
	err error
}

var errStub = errors.New("something wrong")

func (r stubRepository) SaveAndIdentifyLayout(context.Context, string) (string, error) {
	return "", r.err
}

func (r stubRepository) SaveAndIdentifyPerformer(context.Context, string) (string, error) {
	return "", r.err
}

func (r stubRepository) SaveAndIdentifyTask(context.Context, string, string, string) (string, error) {
	return "", r.err
}

func TestNoFiresEventWhenPerformerAddingIsFail(t *testing.T) {
	reg := &Registrar{Repo: stubRepository{err: errStub}}
	_, err := reg.AddPerformer(t.Context(), Performer{})
	if err == nil {
		t.Error("failed try must return error, given nil")
		return
	}
}

func TestFiresEventWhenPerformerAddingIsDone(t *testing.T) {
	reg := &Registrar{Repo: stubRepository{}}
	_, err := reg.AddPerformer(t.Context(), Performer{})
	if err != nil {
		t.Error("failed a performer adding:", err)
		return
	}
	e := reg.Events()
	if len(e) == 0 || e[0] != eventPerformerAdded {
		t.Error("registar no contains events when performer is added, given:", e)
	}
}

func TestFiresEventWhenTaskAddingIsDone(t *testing.T) {
	reg := &Registrar{Repo: stubRepository{}}
	_, err := reg.AddTask(t.Context(), Task{})
	if err != nil {
		t.Error("failed a task adding:", err)
		return
	}
	e := reg.Events()
	if len(e) == 0 || e[0] != eventTaskAdded {
		t.Error("registar no contains events when task is added, given:", e)
	}
}
