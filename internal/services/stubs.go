package services

import (
	"context"

	"timetablerapi/overview"
)

type OverviewRepoStub struct {
}

func (s OverviewRepoStub) FetchPerformerBoard(context.Context, string) (result *overview.PerformerBoard, err error) {
	return overview.NewPerformerBoard("", "", "board"), nil
}
