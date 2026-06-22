package services

import (
	"context"

	"timetablerapi/overview"
	"timetablerapi/registrar"
)

type OverviewStub struct {
}

type RegistrarStub struct {
	Result string
}

func (os OverviewStub) ViewPerformerBoard(context.Context, string) (*overview.PerformerBoard, error) {
	return overview.NewPerformerBoard("", "", "board"), nil
}

func (rs RegistrarStub) AddPerformer(context.Context, registrar.Performer) (string, error) {
	return rs.Result, nil
}

func (rs RegistrarStub) AddTask(context.Context, registrar.Task) (string, error) {
	return rs.Result, nil
}

func (rs RegistrarStub) Events() []uint {
	return nil
}
