package services

import (
	"context"

	"timetablerapi/overview"
)

type OverviewStub struct {
}

type OverviewRepoStub struct {
}

func (s OverviewRepoStub) FetchPerformerBoard(context.Context, string) (result *overview.PerformerBoard, err error) {
	return overview.NewPerformerBoard("", "", "board"), nil
}

func (os OverviewStub) ViewPerformerBoard(context.Context, string) (*overview.PerformerBoard, error) {
	return overview.NewPerformerBoard("", "", "board"), nil
}
