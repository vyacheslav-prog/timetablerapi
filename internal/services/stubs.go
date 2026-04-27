package services

import (
	"context"
	"timetablerapi/overview"
)

type OverviewStub struct {
}

func (os OverviewStub) ViewPerformerBoard(context.Context, string) (*overview.PerformerBoard, error) {
	return overview.NewPerformerBoard("", "", "board"), nil
}
